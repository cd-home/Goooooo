import {ElNotification} from 'element-plus'

const SuccessNotify = (message) => ElNotification({
    title: 'Success',
    message: message,
    type: 'success',
  })

const ErrorNotify = (error) => ElNotification({
    title: 'Error',
    message: error,
    type: 'error',
})

const InfoNotify = (info) => ElNotification({
    title: 'Info',
    message: info,
    type: 'info',
})

const WarnNotify = (warn) => ElNotification({
    title: 'Warning',
    message: warn,
    type: 'warning',
})


const notify = {
    SuccessNotify,
    ErrorNotify,
    InfoNotify,
    WarnNotify
}

export default notify