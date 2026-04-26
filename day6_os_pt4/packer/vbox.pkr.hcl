packer {
  required_plugins {
    virtualbox = {
      version = "~> 1"
      source  = "github.com/hashicorp/virtualbox"
    }
  }
}

variable "pub_key" {
  type    = string
  default = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMGSwLXvWcPv9J/PN9nUPvapyK31T7ShAWNFNzSM1/ry user@arch"
}



source "virtualbox-iso" "vm1" {
  guest_os_type    = "Ubuntu_64"
  iso_url          = "/home/user/Downloads/ubuntu-24.04.4-live-server-amd64.iso"
  iso_checksum     = "sha256:e907d92eeec9df64163a7e454cbc8d7755e8ddc7ed42f99dbc80c40f1a138433"
  vboxmanage = [
      ["modifyvm", "{{.Name}}", "--nat-localhostreachable1", "on"]
  ]
  cpus             = 2
  memory           = 2048
  ssh_username     = "installer"
  ssh_password     = "password"
  ssh_timeout            = "60m"
  ssh_handshake_attempts = 100
  shutdown_command = "echo 'super_secret_password' | sudo -S shutdown -P now"
  boot_wait        = "5s"
  cd_label = "cidata"
  cd_files = [
    "./http/user-data",
    "./http/meta-data",
  ]
  boot_command = [
    "c", "<wait3s>",
    "linux /casper/vmlinuz autoinstall ds=nocloud;s=/cd-rom/", "<enter><wait3s>",
    "initrd /casper/initrd", "<enter><wait3s>",
    "boot", "<enter>"
  ]

}


build {
  name    = "vbox-packer"
  sources = ["sources.virtualbox-iso.vm1"]
  provisioner "shell" {
    inline = [
      "echo -n 'PasswordAuthentication no' | sudo tee -a /etc/ssh/ssh_config",
      "echo -n 'PasswordAuthentication no' | sudo tee -a /etc/ssh/ssh_config",
      "sudo systemctl restart ssh.service",
      "sudo ufw allow OpenSSH",
      "sudo ufw enable",
    ]
  }
}
