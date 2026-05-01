# Ansible Project

This is an Ansible project for automation tasks.

## Structure

- `ansible.cfg`: Ansible configuration
- `inventory/`: Inventory files and variables
- `playbooks/`: Ansible playbooks
- `roles/`: Custom roles
- `collections/`: Ansible collections
- `library/`: Custom modules
- `filter_plugins/`: Custom filter plugins

## Usage

1. Install requirements: `ansible-galaxy install -r requirements.yml`
2. Run a playbook: `ansible-playbook -i inventory/hosts.ini playbooks/playbook.yml`
