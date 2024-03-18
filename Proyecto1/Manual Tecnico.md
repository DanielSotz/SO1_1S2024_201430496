# MANUAL TECNICO PROYECTO 1

## Herramientas de Arquitectura

- ### Docker:
    Plataforma de código abierto que permite la creación, el despliegue y la ejecución de aplicaciones dentro de contenedores. Los contenedores son entornos ligeros y portátiles que contienen todo lo necesario para que una aplicación se ejecute, incluidas bibliotecas, herramientas y código.

- ### Docker Compose:
    Es una herramienta que permite definir y gestionar aplicaciones Docker multi-contenedor de manera fácil y eficiente. Con Docker Compose, puedes definir la estructura de tu aplicación en un archivo YAML
    
- ### DockerHub:
    Es un registro público y privado de imágenes de Docker. Funciona como un repositorio en la nube donde los desarrolladores pueden almacenar, distribuir y compartir imágenes de contenedores Docker.

- ### Volumen de Docker:
    Son mecanismos que permiten persistir datos generados o utilizados por los contenedores Docker. Los volúmenes proporcionan una forma de compartir datos entre contenedores o entre un contenedor y el sistema anfitrión.

- ### Maquina Virtual o VM:
    Es un entorno de computación completo y virtualizado que se ejecuta como un software de aplicación en un sistema operativo físico (llamado "host"). Las máquinas virtuales están aisladas del sistema operativo host y de otras máquinas virtuales en el mismo host, lo que permite ejecutar múltiples sistemas operativos y aplicaciones en un solo hardware físico

- ### VMware: 
    Es un software de virtualización que te permite crear y ejecutar múltiples sistemas operativos en una sola máquina física. 

## Base de Datos

 - ### MySQL:
    Es un sistema de gestión de bases de datos relacional (RDBMS) de código abierto ampliamente utilizado. Se destaca por ser rápido, confiable y fácil de usar. Es compatible con múltiples sistemas operativos y se utiliza en una amplia variedad de aplicaciones.

## Frontend
 - ### React:
    Es una biblioteca de JavaScript de código abierto para construir interfaces de usuario (UI) interactivas y dinámicas. Se utiliza principalmente para el desarrollo de aplicaciones de una sola página (SPA) y aplicaciones de una sola página con múltiples componentes reutilizables. Utiliza un paradigma de programación basado en componentes, lo que significa que las interfaces de usuario se dividen en componentes independientes y reutilizables que se pueden renderizar de manera dinámica en función del estado de la aplicación.

## Backend
 - ### Golang:
     Es un lenguaje de programación de código abierto desarrollado por Google. Se caracteriza por ser rápido, eficiente y fácil de usar. Go se utiliza principalmente para el desarrollo de software de sistemas, herramientas de línea de comandos, servidores web y aplicaciones de redes, entre otros. Go se compila en código máquina nativo, lo que garantiza un rendimiento rápido y eficiente de las aplicaciones desarrolladas con este lenguaje.

## VM
 - ### Ubuntu Server 22.04: 
    Es una versión del sistema operativo Ubuntu diseñada específicamente para servidores. Ubuntu es una distribución de Linux basada en Debian que se ha ganado una amplia adopción en entornos de servidores debido a su estabilidad, seguridad y facilidad de uso.

## Librerias
 - ### Vis.js: 
    es una biblioteca de JavaScript que se utiliza para crear visualizaciones de datos interactivas en la web. Ofrece una amplia variedad de herramientas para crear gráficos de red, gráficos de tiempo, gráficos de árbol, entre otros.
    
 - ### < sys/sysinfo.h>:
    Es parte de la biblioteca estándar de C en sistemas operativos Unix y proporciona funciones para obtener información del sistema, como la cantidad de memoria total, la cantidad de memoria libre, la carga del sistema y otra información relacionada con el sistema.
    
 - ### <linux/mm.h>:
    Proporciona definiciones y funciones relacionadas con la gestión de la memoria en el kernel de Linux. Contiene funciones para la asignación y liberación de memoria, el manejo de páginas de memoria y otras operaciones relacionadas con la gestión de la memoria en el sistema Linux
 - ### <linux/sched.h>:
    Proporciona definiciones y funciones relacionadas con la planificación de tareas (procesos) en el kernel de Linux. Contiene estructuras de datos y funciones para administrar la programación de tareas, la prioridad de las tareas, el estado de las tareas y otras operaciones relacionadas con la planificación de tareas en el sistema Linux.
 - ### <linux/sched/signal.h>:
    Proporciona definiciones y funciones relacionadas con la gestión de señales en el kernel de Linux. Las señales son mecanismos de comunicación entre procesos que se utilizan para notificar eventos o manejar errores. Esta librería contiene estructuras de datos y funciones para la manipulación de señales, la definición de manejadores de señales y otras operaciones relacionadas con las señales en el sistema Linux.

## Otras herramientas y Definiciones
 - ### Lenguaje C:
    Lenguaje de programación de nivel medio que combina la potencia y la flexibilidad del lenguaje ensamblador con la facilidad de uso de los lenguajes de alto nivel. El lenguaje C es ampliamente utilizado en el desarrollo de sistemas operativos, aplicaciones de sistemas embebidos, software de sistemas, y una amplia variedad de aplicaciones de software.

 - ### Kernel:
    Es la parte central de un sistema operativo que controla los recursos del sistema y proporciona servicios a los programas en ejecución. Es responsable de administrar los recursos del hardware, como la memoria, el procesador, los dispositivos de entrada y salida, y las comunicaciones entre los programas y el hardware. 

 - ### Make File:
    Es un archivo de texto especial utilizado por el programa Make para automatizar el proceso de compilación y generación de programas a partir de múltiples archivos fuente. Un Makefile contiene reglas que especifican cómo se deben compilar y vincular los archivos fuente para producir el programa final.

 - ### Kernel Modules:
    Son fragmentos de código que se pueden cargar y descargar dinámicamente en un sistema operativo Linux o Unix mientras el sistema está en ejecución. Los módulos del kernel extienden las capacidades del kernel base al proporcionar controladores de dispositivos adicionales, sistemas de archivos, soporte de red y otras funcionalidades. 

 - ### Proceso:
    Consiste en el código ejecutable del programa, su espacio de memoria, el estado del procesador, la pila de llamadas, los recursos asignados y otra información relacionada con la ejecución del programa. 

 - ### Estados de Procesos:
    Los procesos en un sistema operativo kernel pueden estar en varios estados, que incluyen:
    1. **New**: Cuando se crea un proceso, se almacena en el heap y se llama task.
    2. **Ready**: Aquí se manejan colas de procesos listos para ser ejecutados
    3. **Running**: Es cuando el proceso se ejecuta en el microprocesador.
    4. **Waiting**: Cuando un dispositivo está pidiendo acceso por hardware y cuando ya está listo, regresa a Ready.
    5. **Finished**: Cuando termina un proceso.
- ### Directorio /proc:Directorio /proc
    Contiene un sistema de archivos imaginario o virtual. Se utiliza para ofrecer información relacionada con el sistema
    

