#include <linux/module.h>
#include <linux/moduleparam.h>
#include <linux/init.h>
#include <linux/kernel.h>   
#include <linux/proc_fs.h>
#include <linux/uaccess.h>

#include <linux/utsname.h>
#include <linux/mm.h>

#define BUFSIZE  1000

unsigned long copy_to_user(void __user *to,const void *from, unsigned long n);
unsigned long copy_from_user(void *to,const void __user *from,unsigned long n);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("Liran B.H");


static struct proc_dir_entry *ent;

static ssize_t mywrite(struct file *file, const char __user *ubuf,size_t count, loff_t *ppos) 
{
	printk( KERN_DEBUG "write handler\n");
	return -1;
}

static ssize_t myread(struct file *file, char __user *ubuf,size_t count, loff_t *ppos) 
{

	printk( KERN_DEBUG "read handler\n");

	struct sysinfo i;

	char buf[BUFSIZE];
	int len = 0;

	if(*ppos > 0 || count < BUFSIZE)
		return 0;

	si_meminfo(&i);

	// len += sprintf(buf + len, "%s\n", utsname()->version);
	
	len += sprintf(buf + len, "{");
	len += sprintf(buf + len, "\"Memoria\": %li,\n", i.totalram);
	len += sprintf(buf + len, "\"Libre\": %li,\n", i.freeram);
	len += sprintf(buf + len, "}\n");

	if(copy_to_user(ubuf,buf,len))
		return -EFAULT;
	
	*ppos = len;
	return len;

}

static struct file_operations myops = 
{
	.owner = THIS_MODULE,
	.read = myread,
	.write = mywrite,
};

static int simple_init(void)
{
	printk(KERN_INFO "Hola mundo, somos el grupo 18\n");
	ent=proc_create("mem_grupo18",0660,NULL,&myops);
	return 0;
}

static void simple_cleanup(void)
{
	printk(KERN_INFO "Sayonara mundo, somos el grupo 18\n");
	proc_remove(ent);
}

module_init(simple_init);
module_exit(simple_cleanup);
