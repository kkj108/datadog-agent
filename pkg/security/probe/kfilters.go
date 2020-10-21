// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/DataDog/datadog-agent/pkg/security/rules"
	"github.com/DataDog/datadog-agent/pkg/security/secl/eval"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// ErrDiscarderNotSupported is returned when trying to discover a discarder on a field that doesn't support them
type ErrDiscarderNotSupported struct {
	Field string
}

func (e ErrDiscarderNotSupported) Error() string {
	return fmt.Sprintf("discarder not supported for `%s`", e.Field)
}

// FilterPolicy describes a filtering policy
type FilterPolicy struct {
	Mode  PolicyMode
	Flags PolicyFlag
}

// Bytes returns the binary representation of a FilterPolicy
func (f *FilterPolicy) Bytes() ([]byte, error) {
	return []byte{uint8(f.Mode), uint8(f.Flags)}, nil
}

func isParentPathDiscarder(rs *rules.RuleSet, eventType eval.EventType, filename string) (bool, error) {
	dirname := filepath.Dir(filename)

	bucket := rs.GetBucket(eventType)
	if bucket == nil {
		return false, nil
	}

	event := NewEvent(nil)

	for _, rule := range bucket.GetRules() {
		// ensure we don't push parent discarder if there is another rule relying on the parent path

		// first case: rule contains a filename field
		// ex: rule		open.filename == "/etc/passwd"
		//     discarder /etc/fstab
		// /etc/fstab is a discarder but not the parent

		// second case: rule doesn't contain a filename field but a basename field
		// ex: rule	 	open.basename == "conf.d"
		//     discarder /etc/conf.d/httpd.conf
		// /etc/conf.d/httpd.conf is a discarder but not the parent

		// check filename
		fileField := eventType + ".filename"

		if values := rule.GetFieldValues(fileField); len(values) > 0 {
			for _, value := range values {
				if strings.HasPrefix(value.Value.(string), dirname) {
					return false, nil
				}
			}

			if err := event.SetFieldValue(fileField, dirname); err != nil {
				return false, err
			}

			if isDiscarder, _ := rs.IsDiscarder(event, fileField); isDiscarder {
				return true, nil
			}
		}

		// check basename
		baseField := eventType + ".basename"

		if values := rule.GetFieldValues(baseField); len(values) > 0 {
			if err := event.SetFieldValue(baseField, path.Base(dirname)); err != nil {
				return false, err
			}

			if isDiscarder, _ := rs.IsDiscarder(event, baseField); !isDiscarder {
				return false, nil
			}
		}
	}

	log.Tracef("`%s` discovered as parent discarder", dirname)

	return true, nil
}
