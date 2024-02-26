import React from 'react'

import Navigation from '@/components/navigation/navigation'
// 
const SetUp = () => {
  return (
    <div className='flex justify-center items-center'>
      <Navigation navigation={[]} />
      <div className='bg-red-500 flex-1 py-8 px-56'>
        <h1 className='text-4xl text-slate-950 my-8'>Como iniciar y configurar un proyecto en Node.js</h1>

        <p className='text-2xl text-slate-950 my-8'>En este tutorial nos adentraremos en como crear un proyecto de Node.js paso por paso, desde su instalación hasta configurar nuestro primer servidor.</p>
        <p className='text-2xl text-slate-700 my-8'>¿Listo para arrancar? ¡Arranquemos!</p>

        <p className='text-lg text-slate-950 my-8 my-8'>Node.js es un entorno de tiempo de ejecución de código abierto de JavaScript que permite crear aplicaciones del lado del servidor proporcionando gran rendimiento, eficiencia y escalabilidad a la hora de establecer proyectos</p>

        <p className='text-2xl text-slate-800'>Instalar Node.js</p>
        <p className='text-lg text-slate-800 my-8'>Para empezar la descarga del entorno nos dirigiremos a la <a href="https://nodejs.org/en">página oficial de Node.js (https://nodejs.org/en)</a>
          . Una vez dentro nos encontraremos con dos opciones de descarga: LTS y Current. La diferencia entre ambas es que la primera, la más recomendada, posee la última versión estable y, por el otro lado, la versión Current dispone de las últimas actualizaciones pero no es estable.
          . Para finalizar, la instalación es sencilla, aceptamos los términos y condiciones, elegimos el directorio y, por último, ejecutamos la instalación.</p>

        <p className='text-2xl text-slate-600 my-8'>Comprobar la instalación de Node.js</p>

        <p>Code</p>

        <p className='text-2xl text-slate-600 my-8'>Comprobar la instalación de npm</p>

        <p className='text-lg text-slate-800 my-8'>Una vez alojado el framework, aparte de este mismo, se nos instalará <a href="https://www.npmjs.com/">npm (Node Package Manager)</a>
          que es un gestor de paquetes de Node.js. La misma nos permitira instalar paquetes y librerías de javascript definidas para emplear en nuestros futuros proyectos.</p>

        <p>Code</p>

        <p className='text-2xl text-slate-800 my-8'>Iniciar un proyecto</p>

        <p className='text-lg text-slate-800 my-8'>Para crear un proyecto en Node.js primero creamos una carpeta en el directorio donde desaeamos. Abriremos una consola, ya sea desde el simbolo de sistema del sistema operativo o
          desde de nuestro editor de código, nos dirigiremos a la carpeta donde se encuentra el proyecto y una vez dentro ejecutamos el comando</p>

        <p>Code</p>

        <p className='text-lg text-slate-800 my-8'>¡Hemos inicializado un proyecto con Node.js! Se nos generara un archivo package.json que contendrá la información de nuestro proyecto</p>

      </div>
    </div>
  )
}

export default SetUp