// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import VidulumVidulumVidulum from './vidulum.vidulum.vidulum'


export default { 
  VidulumVidulumVidulum: load(VidulumVidulumVidulum, 'vidulum.vidulum.vidulum'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}