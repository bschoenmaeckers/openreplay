import React from 'react'
import { Icon, Popup } from 'UI'
import stl from './timelinePointer.module.css'
import cn from 'classnames'

function TimelinePointer({ icon, content }) {
  return (
    <Popup
      // offset={20}
      // pinned
      distance={15}
      content={content}
    >
      <div className={cn(stl.wrapper, 'flex items-center justify-center relative')}>          
        <div className={stl.pin} />
        <div style={{ top: '3px' }} className={stl.icon} >
          <Icon name={icon} size="18" style={{ fill: '#D3545F' }} />
        </div>
      </div>
    </Popup>
  )
}

export default TimelinePointer
