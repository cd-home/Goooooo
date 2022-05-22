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

const notify = {
    SuccessNotify,
    ErrorNotify,
}

export default notify