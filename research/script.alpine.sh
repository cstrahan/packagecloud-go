#!/bin/sh

unknown_os ()
{
  echo "Unfortunately, your operating system distribution and version are not supported by this script."
  echo
  echo "You can override the OS detection by setting os= and dist= prior to running this script."
  echo "You can find a list of supported OSes and distributions on our website: https://packagecloud.io/docs#os_distro_version"
  echo
  echo "For example, to force Alpine v3.13: os=alpine dist=v3.13 ./script.sh"
  echo
  echo "Please contact the owner of this repository/package for further support.."
  exit 1
}

curl_check ()
{
  echo "Checking for curl..."
  if command -v curl > /dev/null; then
    echo "Detected curl..."
  else
    echo "Installing curl..."
    apk add curl
    if [ "$?" -ne "0" ]; then
      echo "Unable to install curl! Your base system has a problem; please check your default OS's package repositories because curl should work."
      echo "Repository installation aborted."
      exit 1
    fi
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
}

detect_os ()
{
  if [ -z "${os}" ]; then
    os=alpine
  fi

  if [ -z "${dist}" ]; then
    dist=`/bin/cat /etc/alpine-release | awk -F. '{ print "v"$1"."$2 }'`
  fi

  if [ -z "${dist}" ]; then
    unknown_os
  fi

  # remove whitespace from OS and dist name
  os="${os// /}"
  dist="${dist// /}"

  echo "Detected operating system as $os/$dist."
}

main ()
{
  detect_os
  curl_check

  if [ -z "$unique_id" ]; then
    get_unique_id
  fi

  alpine_config_url="https://REDACTED_TOKEN:@packagecloud.io/install/repositories/USER/REPOSITORY/config_file.alpine?os=${os}&dist=${dist}&name=${unique_id}&source=script"

  echo "Found unique id: ${unique_id}"
  rsa_key_install_url="https://REDACTED_TOKEN:@packagecloud.io/install/repositories/USER/REPOSITORY/rsa_key_url.alpine?os=${os}&dist=${dist}&name=${unique_id}"

  rsa_key_url=`curl -L "${rsa_key_install_url}"`
  if [ "${rsa_key_url}" = "" ]; then
    echo "Unable to retrieve RSA key URL from: ${rsa_key_url}."
    echo "Please contact the owner of this repository/package for further support."
    exit 1
  fi


  alpine_repositories_file="/etc/apk/repositories"
  alpine_keys_dir="/etc/apk/keys"

  echo "Backing up ${alpine_repositories_file} to ${alpine_repositories_file}-${backup_ext}"
  cp ${alpine_repositories_file} ${alpine_repositories_file}-${backup_ext}

  echo -n "Appending to ${alpine_repositories_file}..."

  # create an apt config file for this repository
  curl -sf "${alpine_config_url}" >> ${alpine_repositories_file}
  curl_exit_code=$?

  if [ "$curl_exit_code" = "22" ]; then
    echo
    echo
    echo -n "Unable to download repo config from: "
    echo "${alpine_config_url}"
    echo
    echo "This usually happens if your operating system is not supported by "
    echo "packagecloud.io, or this script's OS detection failed."
    echo
    echo "You can override the OS detection by setting os= and dist= prior to running this script."
    echo "You can find a list of supported OSes and distributions on our website: https://packagecloud.io/docs#os_distro_version"
    echo
    echo "For example, to force Alpine v3.13: os=alpine dist='v3.13'./script.sh"
    echo
    echo "If you are running a supported OS, contact the owner of this repository/package for further support."
    exit 1
  elif [ "$curl_exit_code" = "35" -o "$curl_exit_code" = "60" ]; then
    echo "curl is unable to connect to packagecloud.io over TLS when running: "
    echo "    curl ${alpine_config_url}"
    echo "This is usually due to one of two things:"
    echo
    echo " 1.) Missing CA root certificates (make sure the ca-certificates package is installed)"
    echo " 2.) An old version of libssl. Try upgrading libssl on your system to a more recent version"
    echo
    echo "Contact the owner of this repository/package with information about your system for help."
    exit 1
  elif [ "$curl_exit_code" -gt "0" ]; then
    echo
    echo "Unable to run: "
    echo "    curl ${alpine_config_url}"
    echo
    echo "Double check your curl installation and try again."
    exit 1
  else
    echo "done."
  fi

  alpine_rsakey_full_filename="${alpine_keys_dir}/USER_REPOSITORY.rsa.pub"

  if [ -e $alpine_rsakey_full_filename ]; then
    echo "${alpine_rsakey_full_filename} exists. Backing up to ${alpine_rsakey_full_filename}-${backup_ext}"
    mv ${alpine_rsakey_full_filename} ${alpine_rsakey_full_filename}-${backup_ext}
  fi

  echo "Retrieving RSA key from ${rsa_key_url}, and copying into ${alpine_rsakey_full_filename}"
  curl -fsL "${rsa_key_url}" > ${alpine_rsakey_full_filename}
  
  chmod 0644 ${alpine_rsakey_full_filename}

  echo "done."

  echo
  echo "The repository is setup! You can now install packages."
}

backup_ext=`date '+%Y%m%d-%H%M%S'`
main

