#!/bin/sh

# This is a simple shell script that demonsrtates the use of og-image-generator.
# This script could be run in a git hook to generate og-images

# All files in this folder are markdown files with yaml frontmatter
# The yaml frontmatter contains information like title, description and date
# of the post
posts=`find ~/article_website/src/posts/ -type f`

for post in $posts 
do
  # grep the value of title, description, and date variable from the yaml 
  # frontmatter of this $post file and store in the respective variable
  title=`cat $post | grep title: | awk -F "\"" '{print $2}'`
  slug=`cat $post | grep title: | awk -F "\"" '{print tolower($2)}' | tr " " "_"`
  description=`cat $post | grep description: | awk -F "\"" '{print $2}'`

  # yaml date: field is in format "yyyy-mm-dd" that is formatted by the date
  # util here to a more readable format
  # example 2021-02-28 will be stored as: "28 February 2021"
  date=`cat $post | grep date: | awk -F "\"" '{print $2}' | xargs -I {} date -d {} +"%d %B %Y"`

  # Call og-image-generator with the above variable
  # This requires this command to be in $PATH
  og-image-generator -out="./src/assets/og/$slug.png" -title="$title" -desc="$description" -date="$date"

  echo "==============================================================="
done

