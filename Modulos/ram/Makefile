obj-m := mem_grupo18.o

KERNELDIR ?= /lib/modules/$(shell uname -r)/build

all default: modules
install: modules_install

modules modules_install help clean:
	$(MAKE) -C $(KERNELDIR) M=$(shell pwd) $@


clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
	
test_uno:
	sudo dmesg -C
	sudo insmod mem_grupo18.ko
	sudo cat /proc/mem_grupo18
	sudo rmmod mem_grupo18.ko
	sudo dmesg

