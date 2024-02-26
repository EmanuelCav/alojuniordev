"use client"

import { Pagination, Navigation } from 'swiper/modules';
import { Swiper, SwiperSlide } from 'swiper/react';

import { ITutorial } from '@/interface/Content'

import 'swiper/css';
import 'swiper/css/pagination';
import 'swiper/css/navigation';

import Tutorial from './components/Tutorial';

const ShowTutorials = ({ tutorialSection }: { tutorialSection: ITutorial[] }) => {
    return (
        <Swiper className="flex justify-around items-center flex-row flex-wrap m-8"
            slidesPerView={3}
            centeredSlides={true}
            spaceBetween={30}
            navigation={true}
            style={{
                "--swiper-navigation-color": "rgb(110 231 183)",
              } as any}
            modules={[Pagination, Navigation]}>
            {
                tutorialSection.map((tutorial: ITutorial) => {
                    return (
                        <SwiperSlide key={tutorial.tool}>
                            <Tutorial tutorial={tutorial} />
                        </SwiperSlide>
                    )
                })
            }
        </Swiper>
    )
}

export default ShowTutorials