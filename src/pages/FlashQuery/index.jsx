/* eslint-disable react/jsx-key */
import { Descriptions, Spin } from 'antd'
import { connect } from 'dva'
import moment from 'moment'
import { useIntl } from 'umi'
import { useEffect, useState } from 'react'
import { QueryPageFramework } from '../../components/QueryPageFramework/index.jsx'
import { TableType } from '@/tableType.js'
import styles from '../../global.less';
import { global } from '../../global'

const FlashQuery = ({ dispatch, data, loading }) => {
  const intl = useIntl()
  const [pageSize, setPageSize] = useState(window.innerHeight)

  useEffect(() => {
    if (!global.isFrozen() && global.isFirst()) {
      dispatch({
        type: 'device/frontPage',
        payload: {
          type: TableType.flash,
        }
      })
    }
    window.addEventListener('resize', resizeUpdate)

    return () => {
      window.removeEventListener('resize', resizeUpdate)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [1])

  const resizeUpdate = (e) => {
    setPageSize(e.target.innerHeight)
  }

  const itemCount = 6
  const subtractionHeight = 204
  const descriptionItemHeight = (pageSize - subtractionHeight) / itemCount
  const successCode = 200;

  const stateCardStyle = () => {
    if (data.status === successCode) {
      return styles.stateCodeCardSuccess
    } else {
      return styles.stateCodeCardFail
    }
  }

  return (
    <QueryPageFramework
      dispatch={dispatch}
      loading={loading}
      type={TableType.flash}
    >
      <div className={stateCardStyle()}>
        {
          data.status === successCode ?
            intl.formatMessage({ id: "pages.table.requestSuccess" }) :
            intl.formatMessage({ id: "pages.table.requestFailureCode" }) +
            data.errorCode
        }
      </div>
      <Descriptions bordered className={styles.table} style={{ textAlign: 'center' }}>
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.detectionTime" })} span={4}>
          {data.detectionTime == "N/A" ? data.detectionTime : moment(data.detectionTime).format('YYYY-MM-DD hh:mm')}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.sampleSerialNumber" })} span={4}>
          {data.sampleNumber}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.atmosphericPressure" })} span={4}>
          {data.pressure}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.preFlashTemperature" })} span={4}>
          {data.preTemperature}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.flashPointTemperature" })} span={4}>
          {data.pointTemperature}
        </Descriptions.Item >
      </Descriptions>
    </QueryPageFramework >
  )
}

export default connect(({ device, loading }) => ({
  data: device,
  loading: loading.models.device,
}))(FlashQuery)