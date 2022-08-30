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

const AcidQuery = ({ dispatch, data, loading }) => {
  const intl = useIntl()
  const [pageSize, setPageSize] = useState(window.innerHeight)

  useEffect(() => {
    if (!global.isFrozen() && global.isFirst()) {
      dispatch({
        type: 'device/frontPage',
        payload: {
          type: TableType.acid,
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

  const itemCount = 5.5
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
    (data.items ?
      <QueryPageFramework
        dispatch={dispatch}
        loading={loading}
        type={TableType.acid}
      >
        <div className={stateCardStyle()}>
          {
            data.status === successCode ?
              intl.formatMessage({ id: "pages.table.requestSuccess" }) :
              intl.formatMessage({ id: "pages.table.requestFailureCode" }) +
              data.errorCode
          }
        </div>
        <Descriptions bordered className={styles.table}>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={intl.formatMessage({ id: "pages.table.detectionTime" })} span={4}>
            {data.detectionTime == 'N/A' ? data.detectionTime
              : moment(data.detectionTime).format('YYYY-MM-DD hh:mm')}
          </Descriptions.Item >
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={intl.formatMessage({ id: "pages.table.number" })} span={4}>
            {data.number}
          </Descriptions.Item >
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 1} span={2}>
            {data.items[0]}
          </Descriptions.Item>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 2} span={2}>
            {data.items[1]}
          </Descriptions.Item>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 3} span={2}>
            {data.items[2]}
          </Descriptions.Item>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 4} span={2}>
            {data.items[3]}
          </Descriptions.Item>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 5} span={2}>
            {data.items[4]}
          </Descriptions.Item>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={"PH" + 6} span={2}>
            {data.items[5]}
          </Descriptions.Item>
        </Descriptions>
      </QueryPageFramework > : <Spin size="large" />)
  )
}

export default connect(({ device, loading }) => ({
  data: device,
  loading: loading.models.device,
}))(AcidQuery)