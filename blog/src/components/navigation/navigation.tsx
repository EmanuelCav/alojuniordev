import React from 'react'

import { NavigationTypeProps } from '@/types/props.types'

const Navigation = ({ navigation }: { navigation: NavigationTypeProps[] }) => {
  return (
    <div className='w-1/4 p-8 bg-blue-500'>Navigation</div>
  )
}

export default Navigation