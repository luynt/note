### Issues:

##### Issue 1: *grub error: you need to load kernel first**
- Find partition contain boot
```bash
grub> ls
    (hd0) (hd0,gpt7) (hd0,gpt6) (hd0,gpt5) (hd0,gpt4) (hd0,gpt3) (hd0,gpt2) (hd0,gpt1)
```
Assume partition *(hd0,gpt1)* contain boot info
- Set root envirionment
```bash
grub> set root=(hd0,gpt1)
```
- Set kernel version
```bash
grub> linux /boot/vmlinux-<specific-version>  root=/dev/sda1
```
- Init 
```bash
grub> initrd /boot/initrd.img-<specific-version>
```
- Reboot PC
```bash
grub> boot
```
