# (Updating) Server manipulation basics for beginners: virtual machine, ssh connection and Ansible

## To begin...

This note is to keep track of my progress learning the basics of server manipulation during my summer as an IT intern for [M-Service](https://momo.vn/). To be able to write this, I want to say thank you, in general, to the whole infrastructure team, the company's committee and HR team for giving me a good first-hand experience in the fintech industry and taking good care of me, and specially to Vinh Vo, who acted as a colleague and a mentor during my internship, and hope that you will find a good girlfriend that suits your enthusiasm well in the future. I hope that you, who are reading this note (probably a new intern guy at Momo), will have fun and enjoy the ride!

## Topics
[Introduction](#intro)
  * [Background](#bkg)
  * [Setting up virtual machine](#vm)
  * [Basic Linux commands](#basic)
  * [Setting up ssh connection](#ssh)
  
[Ansible](#ans)

## <a name="intro"></a> Introduction

### <a name="bkg"></a> Background

By the time I had my internship, I was a beginner to computing, only knowing some basic C and Java, and did not know how to operate on linux terminals. I used Windows as my operating system back then, and my knowledge about systems and servers are plain zero. Unlike what I expected, despite being related, this field does not require prereq coding knowledge, so I belive that as long as you enjoy it, you will get the hang of it ;) .

A bit about my job: The company I worked for, M-Service, is well-known in Vietnam for its online payment application Momo. I worked here in the technology department, in the infrastructure team as a DevOps trainee.

### <a name="vm"></a> Setting up virtual machine

As said above, I used Windows as my operating system. So hectic when it comes to operating systems. It has so many limitations (many operations does not support Windows) that every time I googled to find a solution, another one popped up almost immediately. So I decided to have another machine that operates on ubuntu. There are many ways to solve, imo:

_Buy another PC (not enough money :( )

_Dual boot (you may lose data if you made a mistake)

_**Setting up virtual machine**: I chose this method!

#### How:
To set up a virtual machine, I used VMWare. See this video:

<a href="http://www.youtube.com/watch?feature=player_embedded&v=BHpRTVP8upg
" target="_blank"><img src="http://img.youtube.com/vi/BHpRTVP8upg/0.jpg" 
alt="thumbnail" width="240" height="180" border="10" /></a>

.iso disc file can be found on [Ubuntu's main download page](https://ubuntu.com/download). I would recommend using Ubuntu Server as a lightweight choice.

### <a name="basic"></a> Basic Linux commands, if you don't know yet

* Example of a directory: `/etc/ansible/file-name`

* `cd directory`: change working directory.

* `ls directory`: show what files and folder are in that directory. If no directory is specified, info will be shown on the working directory.

* `cp ...`: copying files. See [this instruction](https://shapeshed.com/unix-cp/).

* `sudo op_name command`: perform command using an installed operation.

* Editing a file: `sudo editor file_directory`. For `editor`, I would recommend `nano` and `vim`. By default they are installed with the OS, and `nano` is easier to use if you know nothing yet.

### <a name="ssh"></a> Setting up ssh connection

#### What you need to know

When connecting, a machine will have:
* Connection type (which is ssh)
* A username (in the form of xxx@xxx)
* An ip address
* A connecting port (22 by default, ranged from 0 to 2^16-1)

When connected, the connecting machince will be able to perform actions on behalf of the connected.

#### Windows

As far as I have known, Windows can only connect to other servers, and is not able to be connected.

To connect, download [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html). The rest is just opening it up and fill in the ip address and the port.

#### Ubuntu
Note: any error that is related to permission and access, see if you have typed `sudo` before any command.

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

## <a name="ans"></a> Ansible: common occuring problems

Ansible has already have its own document guide. For basic learning, you can visit [Ansible User Guide](https://docs.ansible.com/ansible/latest/user_guide/index.html).

Here, I will just note down on problems upon learning Ansible that I have approached and solved using external search outside the document. Treat this as a troubleshooting checklist for beginners to save you time.

Problem type | Why? | Solution
--- | --- | ---
`host1 UNREACHABLE!` | By default, the system will use the username `root` while your username is not `root`. | [ansible_user](#edit) or add `--become-user <username>` flag to the command 
Same as above | The server to connect to has password, and you have not set it to plug the password in. | [ansible_password](#edit), or add `--ask-password` flag to the command to manually type in password.

Solution includes:

### <a name="edit"></a>Editing inventory files:

Add the specified keyword into your inventory file, after the server name/ip address (default is `/etc/ansible/hosts`), using a text editor. See [this keyword list and how to add them](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html#connecting-to-hosts-behavioral-inventory-parameters)

Example in host file:

```
#Use the host file from the username taymonkhanh when connecting to servername.com
servername.com ansible_user=taymonkhanh 

#Connect to port 69 instead of the default (22) when connecting to server floorgreen.us
floorgreen.us ansible_port=69 
```
