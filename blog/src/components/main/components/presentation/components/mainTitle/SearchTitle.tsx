import React from 'react'
import SearchTagTitle from './components/SearchTagTitle'

const SearchTitle = () => {
    return (
        <div className='flex-1 h-full px-16 flex justify-evenly items-center flex-col'>
            <SearchTagTitle text='Encuentra planes y ejemplos prácticos con el fin de llegar a los mejores resultados' title='Proyectos' />
            <SearchTagTitle text='Descubre guías para comenzar a manejar una herramienta deseada' title='Tutoriales' />
        </div>
    )
}

export default SearchTitle