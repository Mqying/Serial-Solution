import { login } from '@/services/login/index'
import { message, Button, Image } from 'antd'
import { history, useIntl } from 'umi'
import styles from './index.less'

const Login = () => {
  const intl = useIntl()
  const successCode = 200

  const handleSubmit = async () => {
    const msg = await login();
    if (msg.status === successCode) {
      const defaultLoginSuccessMessage = intl.formatMessage({
        id: 'pages.login.success',
        defaultMessage: 'login success',
      })

      message.success(defaultLoginSuccessMessage)

      if (!history) return
      const { query } = history.location
      const { redirect } = query
      history.push(redirect || '/')
      return
    } else {
      const defaultLoginFailureMessage = intl.formatMessage({
        id: 'pages.login.failure',
        defaultMessage: 'login failure!',
      })

      message.error(defaultLoginFailureMessage)
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div>
          <Image
            preview={false}
            src="login.png"
          />
        </div>
        <h1 className={styles.text}>{
          intl.formatMessage({
            id: 'pages.login.title'
          })
        }</h1>
        <div className={styles.buttonText}>
          <Button type="primary" block onClick={() => handleSubmit()}>
            {
              intl.formatMessage({
                id: 'pages.login.button'
              })
            }
          </Button>
        </div>
      </div>
    </div >
  )
}

export default Login
