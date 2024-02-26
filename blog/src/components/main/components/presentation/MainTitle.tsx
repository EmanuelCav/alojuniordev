
import Title from './components/mainTitle/Title'
import SearchTitle from './components/mainTitle/SearchTitle'

const MainTitle = () => {
  return (
    <div className='h-1/2 w-full flex justify-center items-center flex-row'>
      <Title />
      <SearchTitle />
    </div>
  )
}

export default MainTitle