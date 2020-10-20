#!/usr/bin/env bash
set -euo pipefail

printf '=%.0s' {0..79} ; echo
set -x

cd "$(dirname "$0")"

git clean -fdx .

# Generate ssh-key and ignition files
./01-ignition.sh
IGNITION_BASE64=$(base64 -w 0 ignition.json)

REGION="${REGION:-us-east-1}"
UPDATE_STREAM="${UPDATE_STREAM:-stable}"
AMI="$(curl "https://builds.coreos.fedoraproject.org/streams/${UPDATE_STREAM}.json" | jq -r ".architectures.x86_64.images.aws.regions.\"$REGION\".image")"

# TODO remove the IamInstanceProfile
tee specification.json << EOF
{
  "ImageId": "${AMI}",
  "InstanceType": "t2.medium",
  "Monitoring": {
    "Enabled": false
  },
  "BlockDeviceMappings": [
    {
      "DeviceName": "/dev/xvda",
      "Ebs": {
        "DeleteOnTermination": true,
        "VolumeSize": 15,
        "VolumeType": "gp2"
      }
    }
  ],
  "UserData": "${IGNITION_BASE64}",

  "SubnetId": "subnet-c18341ed",
  "IamInstanceProfile": {
    "Name": "ci-datadog-agent-e2e-runner"
  },
  "SecurityGroupIds": ["sg-0f5617ceb3e5a6c39"]
}
EOF

echo "Using DATADOG_AGENT_IMAGE=${DATADOG_AGENT_IMAGE}"
echo "Running inside a gitlab pipeline, using DATADOG_AGENT_IMAGE=${DATADOG_AGENT_IMAGE}"

# Check is the image is hosted on a docker registry and if it's available
if [[ "${DATADOG_AGENT_IMAGE%/*}" == "datadog" ]]
then
    echo "${DATADOG_AGENT_IMAGE} is hosted on a docker registry, checking if it's available"
    IMAGE_REPOSITORY=${DATADOG_AGENT_IMAGE%:*}
    IMAGE_TAG=${DATADOG_AGENT_IMAGE#*:}
    curl -Lfs "https://registry.hub.docker.com/v1/repositories/${IMAGE_REPOSITORY}/tags" | \
        jq -re ".[] | select(.name==\"${IMAGE_TAG}\")" || {
            echo "The DATADOG_AGENT_IMAGE=${DATADOG_AGENT_IMAGE} returns a 404 on the registry.hub.docker.com"
            exit 2
    }
fi

exec ./02-ec2.sh
