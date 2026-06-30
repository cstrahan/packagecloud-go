#!/bin/bash

curl_check ()
{
  echo "Checking for curl..."
  if command -v curl > /dev/null; then
    echo "Detected curl..."
  else
    echo "curl could not be found, please install curl and try again"
    exit 1
  fi
}

get_unique_id ()
{
  echo "A unique ID was not specified, using the machine's hostname..."

  unique_id=`hostname -f 2>/dev/null`
  if [ "$unique_id" = "" ]; then
    unique_id=`hostname 2>/dev/null`
    if [ "$unique_id" = "" ]; then
      unique_id=$HOSTNAME
    fi
  fi

  if [ "$unique_id" = "" -o "$unique_id" = "(none)" ]; then
    echo "This script tries to use your machine's hostname as a unique ID by"
    echo "default, however, this script was not able to determine your "
    echo "hostname!"
    echo
    echo "You can override this by setting 'unique_id' to any unique "
    echo "identifier (hostname, shasum of hostname, "
    echo "etc) prior to running this script."
    echo
    echo
    echo "If you'd like to use your hostname, please consult the documentation "
    echo "for your system. The files you need to modify to do this vary "
    echo "between Linux distribution and version."
    echo
    echo
    exit 1
  fi
  token=`curl -s -XPOST --data "name=${unique_id}" https://REDACTED_TOKEN:@packagecloud.io/install/repositories/USER/REPOSITORY/tokens.text`

}

virtualenv_check ()
{
  if [ -z "$VIRTUAL_ENV" ]; then
    echo "Detected VirtualEnv: Please visit https://packagecloud.io/USER/REPOSITORY/install#virtualenv"
  fi
}

new_global_section ()
{
  echo "No pip.conf found, creating"
  mkdir -p "$HOME/.pip"
  pip_extra_url > $HOME/.pip/pip.conf
}

edit_global_section ()
{
  echo "pip.conf found, making a backup copy, and appending"
  cp $HOME/.pip/pip.conf $HOME/.pip/pip.conf.bak
  awk -v regex="$(escaped_pip_extra_url)" '{ gsub(/^\[global\]$/, regex); print }' $HOME/.pip/pip.conf.bak > $HOME/.pip/pip.conf
  echo "pip.conf appended, backup copy: $HOME/.pip/pip.conf.bak"
}

pip_check ()
{
  version=`pip --version`
  echo $version
}

abort_already_configured ()
{
  if [ -e "$HOME/.pip/pip.conf" ]; then
    if grep -q "USER/REPOSITORY" "$HOME/.pip/pip.conf"; then
      echo "Already configured pip for this repository, skipping"
      exit 0
    fi
  fi
}

pip_extra_url ()
{
  printf "[global]\nextra-index-url=https://${token}:@packagecloud.io/USER/REPOSITORY/pypi/simple\n"
}

escaped_pip_extra_url ()
{
  printf "[global]\\\nextra-index-url=https://${token}:@packagecloud.io/USER/REPOSITORY/pypi/simple"
}

edit_pip_config ()
{
  if [ -e "$HOME/.pip/pip.conf" ]; then
    edit_global_section
  else
    new_global_section
  fi
}

main ()
{
  abort_already_configured
  curl_check
  virtualenv_check
  pip_check

  if [ -z "$unique_id" ]; then
    get_unique_id
  fi
  echo $unique_id

  edit_pip_config

  echo "The repository is setup! You can now install packages."
}

main

