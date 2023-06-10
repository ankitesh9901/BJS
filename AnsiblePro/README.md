# Ansible Script to Install Docker and Setuptools

This Ansible script automates the installation of Docker and Setuptools on the specified hosts using a dynamic inventory.

## Prerequisites

- Ansible: Make sure Ansible is installed on the system where you'll be running this script.

## Usage

1. Put aws_ec2.yaml at /opt/ansible/inventory/ and put your AWS key and secret inside it.

2. Put ansible.cfg at /etc/ansible/ansible.cfg


Make sure you are in the same directory where the playbook file is located.

## Playbook Description

- The playbook installs Docker and Setuptools on the specified hosts.
- It uses the `dnf` module to manage package installation.
- The `ansible_python_interpreter` variable is set to `/usr/bin/python3.9` to ensure Ansible uses the correct Python interpreter for the tasks.
- The playbook also starts the Docker daemon using the `service` module.

## License

This project is licensed under the [MIT License](LICENSE).
