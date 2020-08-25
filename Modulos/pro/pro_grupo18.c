/*


Posterior a esto deberá listar todos los procesos, mostrando:

PID
Nombre
Estado



*/
#include <linux/module.h> 
#include <linux/kernel.h> 
#include <linux/init.h>
#include <linux/list.h>
#include <linux/types.h>
#include <linux/slab.h>
#include <linux/sched.h>
#include <linux/string.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h> 
#include <linux/hugetlb.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>

#define FileProc "pro_grupo18"
#define Carnet ""
#define Nombre "G18"
#define Curso "Sistemas operativos 2"
#define SO "Ubuntu"

struct task_struct *task;
struct task_struct *task_child;
struct list_head *list;
long int ejecucion;
long int suspendido;
long int detenido;
long int zombie;

static int proc_llenar_archivo(struct seq_file *m, void *v) {

	ejecucion = 0;
	suspendido = 0;
	detenido = 0;
	zombie = 0;

  

	for_each_process(task){

		if(task->state == 0){

			ejecucion = ejecucion +1;

		}else if(task->state == 1){
			suspendido = suspendido + 1;

		}else if(task->state == 128){

			detenido = detenido + 1;

		}else if(task->state == 260){

			zombie = zombie + 1;
		}



	
		list_for_each(list, &task->children){

			

			task_child = list_entry(list, struct task_struct, sibling);

			if(task_child->state == 0){

				ejecucion = ejecucion +1;

			}else if(task_child->state == 1){
				suspendido = suspendido + 1;

			}else if(task_child->state == 128){

				detenido = detenido + 1;

			}else if(task_child->state == 260){

				zombie = zombie + 1;
			}



		}
			
		

		
	}


    
	seq_printf(m,"{\"Ejecucion\": %li, \"Suspendido\": %li, \"Detenido\": %li, \"Zombie\":%li}\n", ejecucion, suspendido, detenido, zombie);
    return 0;
}



static int proc_al_abrir_archivo(struct inode *inode, struct  file *file) {
  return single_open(file, proc_llenar_archivo, NULL);
}

static struct file_operations myops =
{
        .owner = THIS_MODULE,
        .open = proc_al_abrir_archivo,
        .read = seq_read,
        .llseek = seq_lseek,
        .release = single_release,
};



static int simple_init(void){

    proc_create(FileProc,0,NULL,&myops);
    printk(KERN_INFO "Hola mundo, somos el grupo 18 - proc %s\n", Nombre);
    return 0;
}

static void simple_clean(void){

    printk(KERN_INFO "“Sayonara mundo, somos el grupo 18\n");
    remove_proc_entry(FileProc,NULL);
}



module_init(simple_init);
module_exit(simple_clean);
/*
 * Documentacion del modulo
 */
MODULE_LICENSE("GPL");
MODULE_AUTHOR(Nombre);
MODULE_DESCRIPTION("Modulo para mostrar info de los proc");
