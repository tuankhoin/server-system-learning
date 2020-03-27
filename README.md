# Internship note: Server learning

## To begin...

This note is to keep track of my one-month progress learning the basics of server manipulation during my summer as an IT intern for Momo, or [M-Service](https://momo.vn/). To be able to write this, I want to say thank you, in general, to the whole infrastructure team, the company's committee and HR team for giving me a good first-hand experience in the fintech industry and taking good care of me, and specially to Vinh Vo, who acted as a colleague and a mentor during my internship, and hope that you will find a good girlfriend that suits your enthusiasm well in the future. I hope that you, who are reading this note, will have fun and enjoy the ride!

Who I am writing this for: a new trainee at Momo, or a server developing enthusiast with little or none of background knowledge in Linux and CLI.

## Topics
* [Introduction](#intro)
  * [Background](#bkg)
  * [Setting up virtual machine](#vm)
  * [Basic Linux commands](#basic)
  * [Setting up ssh connection](#ssh)
* [Ansible](#ans)
* [Modifying files remotely using sftp connection and FileZilla](#sftp)
* [Golang](#go)
* [More?](#what)

## <a name="intro"></a> Introduction

### <a name="bkg"></a> Background

By the time I had my internship, I was a beginner to computing, only knowing some basic C and Java, and did not know how to operate on linux terminals. I only knew Windows when using operating systems back then, and my knowledge about systems and servers were plain zero. Unlike what I expected, despite being related, this field does not require too much prereq coding knowledge, so I belive that as long as you enjoy it, you will mostly get the hang of it .

A bit about my job: The company I worked for, M-Service, is well-known in Vietnam for its online payment application Momo. I worked here in the technology department, in the infrastructure team as a DevOps trainee.

### <a name="vm"></a> Setting up virtual machine

As said above, I used Windows as my operating system. A regrettable choice. It has so many limitations (many operations does not support Windows) that every time I googled to find a solution, another problem popped up almost immediately. So I decided to have another operating system, using Ubuntu. There are many ways to do this:

_Buying another PC

_Dual booting (you may lose data if you make a mistake)

_Use tools like Cygwin or MinGW.

_Setting up virtual machine: This is my choice.

#### How:
To set up virtual machines, I used VMWare. See this video on how to set up one:

<a href="http://www.youtube.com/watch?feature=player_embedded&v=BHpRTVP8upg
" target="_blank"><img src="http://img.youtube.com/vi/BHpRTVP8upg/0.jpg" 
alt="thumbnail" width="480" height="360" border="10" /></a>

Recommended settings for each virtual machine:
* 15-20GB for hard disk.
* The main server to operate on, connect and control the remote ones: 4GB RAM, 2*2 processors
* The remote servers, just to test your connection and deployments: 2GB RAM, 1*1 processors

.iso disc file can be found on [Ubuntu's main download page](https://ubuntu.com/download). I would recommend using Ubuntu Server (only Command Line, no UI) as a lightweight choice.

**Try this:** Set up 3 virtual machines: a main virtual machine as a control node, and 1 to 2 remote machines to operate on.

### <a name="basic"></a> Basic Linux commands, if you don't know yet

* Example of a directory: `/etc/ansible/file-name`

* `cd directory`: change working directory.

* `ls directory`: show what files and folder are in that directory. If no directory is specified, info will be shown on the working directory.

* `cp ...`: copying files. See [this instruction](https://shapeshed.com/unix-cp/). See also, with caution: `mv`, `rm`.

* `sudo op_name command`: perform command using an installed operation. `sudo` acts as a giver of superuser access to bring permission on operations.

* Editing a file: `sudo editor file_directory`. For `editor`, I would recommend `nano` and `vim`. By default they are installed with the OS. `nano` is easier to use if you know nothing yet, and `vim` is more versatile when you are ready to learn.

* `mkdir location/directory_name/`: Creating a new directory/folder.

* Scrolling up or down the terminal: `Shift` + `Page up`/`Page down`.

* `su`: login as root.

**Tips:** During a manual installation I have accidently made a mistake: moving the whole system folder `/etc` to a non-existent location and ended up losing a day to install a new virtual machine and bring the old files in.

* How to avoid: See this [thread](https://askubuntu.com/questions/841449/file-lost-using-mv).
* How to access your disk files: the .vmdk files, your hard drive files (you can find them in your set Virtual Machine storing folder) can be opened as an archive using 7-Zip. You can extract most of the files, but you cannot modify it.

### <a name="ssh"></a> Setting up ssh connection

SSH stands for Secure Shell. It lets you operate a server remotely and securely.

#### What you need to know

When connecting, a machine will have:
* Connection type (which is ssh)
* A username (in the form of xxx@xxx)
* An ip address
* A connecting port (22 by default, ranged from 0 to 2^16-1)

When connected with ssh, the connecting machince will be able to perform actions on behalf of the connected.

#### Windows

As far as I have known, Windows can only connect to other servers, and is not able to be connected.

To connect, download [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html). The rest is just opening it up and fill in the ip address and the port.

#### Ubuntu
Note: any error that is related to permission and access, see if you have typed `sudo` before any command. If not, see [this thread](https://askubuntu.com/questions/230476/how-to-solve-permission-denied-when-using-sudo-with-redirection-in-bash)

Normally Ubuntu Server would give you the choice of installing OpenSSH from installation already. If not:

```
$ sudo apt update
$ sudo apt install openssh-server
```

Now you are ready! 

##### The connected:

Once ssh is enabled, the machine will be able to be connected.

There are many ways to get the ip address, such as `ip a` or `ifconfig`. The address will appear after the word `inet` in the form of xxx.xxx.xxx.xxx (`192.168.142.122` for example).

For configurations such as changing port or setting keys, see [Ubuntu guide](https://help.ubuntu.com/lts/serverguide/openssh-server.html)

##### The connect:

```
$ ssh ip_address_of_the_connected
```

Typing `yes` when asked, doing some login steps and you'll be ok in most cases.

For more, see [this detailed guide](https://linuxize.com/post/how-to-enable-ssh-on-ubuntu-18-04/)

**Try this:** Connect a Linux operated (probably your main virtual machine), then a Windows operated (with PuTTY), to the 2 remote virtual machines (`logout` when you want to exit a machine).

## <a name="ans"></a> Ansible: common occuring problems

Ansible is useful for managing remmote servers, and perform actions on them, such as rebooting or deploying packages. It has 2 types of controlling:

* Ad-hoc command: direct command to perform an action.
* Playbook: perform a series of commands on chosen servers. Very useful for repeated tasks.

Ansible has already have its own document guide. For basic learning, you can visit [Ansible User Guide](https://docs.ansible.com/ansible/latest/user_guide/index.html). For a more simplified instruction, see [this tutorial](https://serversforhackers.com/c/an-ansible-tutorial).

Here, I will just note down on problems upon learning Ansible that I have approached and solved using external search outside the document. Treat this as a troubleshooting checklist for beginners to save your time.

**Try this:** Write an Ansible playbook that installs Docker Engine to the remote machines. Ad-hoc instructions can be found [here](https://docs.docker.com/install/linux/docker-ce/ubuntu/). If you don't know anything during the process, searching them on the internet is relatively easy. Trust me!

**Tips**: 
To put commands of other packages into the playbook (such as apt or docker), Ansible docs has a documentation on how to convert them. These might be useful for you: [apt](https://docs.ansible.com/ansible/latest/modules/apt_module.html), [apt_key](https://docs.ansible.com/ansible/latest/modules/apt_key_module.html), [apt_repository](https://docs.ansible.com/ansible/latest/modules/apt_repository_module.html).

Problem type | Why? | Solution
--- | --- | ---
`host1 UNREACHABLE!` | By default, the system will use the username `root` while your username is not `root`. | [ansible_user](#edit) or add `--become-user <username>` flag to the command 
Same as above | The server to connect to has password, and you have not set it to plug the password in. | [ansible_password](#edit), or add `--ask-password`/`--ask-become-password` flag to the command to manually type in password.
All .yaml related errors | They are mostly syntax errors. Your writing is probably not legal YAML and needs to check up. | [yamllint](#lint)
`Authentication or permission failure. In some cases,...` | By default your tmp directory is probably put elsewhere. | Modify `ansible.cfg`. [Source](https://github.com/ansible/ansible/issues/43830#issuecomment-450808102)

Solution includes:

### <a name="edit"></a>Editing inventory files:

Add the specified keyword into your inventory file, after the server name/ip address (default is `/etc/ansible/hosts`), using a text editor. See [this keyword list and how to add them](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html#connecting-to-hosts-behavioral-inventory-parameters)

Example in host (`.ini` format) file:

```
#When connecting to servername.com:

#Use the host file from the username taymonkhanh
servername.com ansible_user=taymonkhanh 

#Connect to port 69 instead of the default (22)
servername.com ansible_port=69 
```

### <a name="lint"></a>.yaml syntax error handling

Install yamllint, and use this to check if your file is legal `.yaml`. It will appear error locations for you as well.

Note: Some errors may occur because of the accidential use of underscores and dashes. A `_` may be mistakenly become `-` and vice versa.

## <a name="sftp"></a> sftp connect using FileZilla

SFTP stands for SSH File Transfer Protocol. Just like its name, SFTP connects using SSH to transer files between servers. For a simple use of it, try FileZilla to get/upload files between servers.

Step 1: Install [FileZilla](https://filezilla-project.org/)

Step 2: Provide sudo access to your logging account with `chown -R username directory`

Step 3: Now you can config, retrieve and add files to your connected server.

## <a name="go"></a>Golang

Golang is a programming language written by Google for back-end servers. Would be nice to learn if you are into concurrency. Background knowledge of C is recommended IMHO. 

It took me 3 days full to learn all the fundamentals from scratch on [A Tour of Go](https://tour.golang.org/list) (concurrency took me a whole day) and briefly compose the [basic-review.go](https://github.com/tuankhoin/server-system-learning/blob/master/basic-review.go), so I think that you will get used to it fairly fast given a fair background (I only know Matlab, C and Java by the time I'm writing this fyi).

* [Install Go](https://golang.org/doc/install)
* [Setting up and run your first program](https://golang.org/doc/install?download=go1.13.7.linux-amd64.tar.gz)
* Learning Go: [Use A Tour of Go as your playground guide](https://tour.golang.org/list)
* To compile and test your code online, try [Golang Playground](https://play.golang.org/)

### Error handling
* Path resets after reboot: You may have typed the path export to the command line instead of adding it to profile.
  * Method 1: Adding the command to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile` as the documentation says.
  * Method 2: Either adding the path export command to the local bashrc `~/.bashrc` (if applying on 1 user only) or `/etc/bash.bashrc` (if applying globally)
  * See more: [Difference between .bashrc and .profile](https://www.pixelstech.net/article/1478399975-Differences-between-bashrc-and-profile-in-Linux)
* `Command not found`: make sure the `$PATH` is set correctly. Read above again.
* `Permission denied` and cannot use `sudo`: try `sudo chmod -R 777 ~/go`, where the `~/go` is your project directory (`go` is the default name in the instruction. You can set it to your preference). [Source](https://www.reddit.com/r/golang/comments/3ho293/permission_denied_when_using_go_install/)

## <a name="what"></a>More stuff?

Here I will put a list of remaining interesting tools that I could not finished, due to either the limited internship time or lack of backgound knowledge. All of these will be very interesting with many applications once you are proficient:

### [Terraform](https://www.terraform.io/)
Background knowledge: Cloud services

Terraform works like Ansible, deploying tasks to the remote servers (Infrastructure ad Code). The difference is that Terraform is declarative. It runs on cloud servers only, so you will need to have knowledge on at least one cloud provider.

### [Odoo for Developers](https://www.odoo.com/documentation/13.0/index.html)
Background knowledge: Basic Python

Odoo has various purposes. A front-end feature that you can try on Odoo for Developers is to design a UI. Python will be used for development.

### [Docker](docker.io)

You already know virtual machines right from the first part. But it might be very resource consuming to run simultaneous applications on them. Container is a more lightweight, movable solution, and Docker is to build, test and deploy them. [How is it more lightweight?](https://www.quora.com/What-is-the-purpose-of-Docker/answer/Anjali-Nair-324)

### [Git using CLI](https://git-scm.com/book/en/v2/Getting-Started-The-Command-Line)

To be more fluent with CLI, you can try modifying repositories via git, but using only CLI.
