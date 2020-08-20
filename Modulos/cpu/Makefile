obj-m := cpu_grupo18.o

KERNELDIR ?= /lib/modules/$(shell uname -r)/build

all default: modules
install: modules_install

modules modules_install help clean:
	$(MAKE) -C $(KERNELDIR) M=$(shell pwd) $@
insert:
	sudo insmod cpu_grupo18.ko
	sudo dmesg
	sudo cat /proc/cpu_grupo18

desinsert:
	sudo rmmod cpu_grupo18.ko
	sudo dmesg	


clean:

	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
