#!/bin/sh

export FILENAME_IMAGES_TO_SCAN=$1
export REPORTS_DIRECTORY="tmp-scanning-reports"
export FILENAME_OUTPUT=$2

set -x

mkdir -p $REPORTS_DIRECTORY 
cat $FILENAME_IMAGES_TO_SCAN | xargs -I {} sh -c './clair-scanner -c http://docker:6060 --ip $(hostname -i) -r "$REPORTS_DIRECTORY"/report-"$RANDOM".json -l clair.log -w clair-whitelist.yml {}'
ls $REPORTS_DIRECTORY | xargs -I % sh -c "cat $REPORTS_DIRECTORY/% | jq '.image as \$prefix | .vulnerabilities[].description |= \$prefix + \" - \" + .' > $REPORTS_DIRECTORY/processed-%"
jq -s '[.[] | to_entries] | flatten | reduce .[] as $dot ({}; .[$dot.key] += $dot.value) | .image = "nlxio" | .unapproved |= unique' $REPORTS_DIRECTORY/processed-* > $FILENAME_OUTPUT 
rm -rf $REPORTS_DIRECTORY
