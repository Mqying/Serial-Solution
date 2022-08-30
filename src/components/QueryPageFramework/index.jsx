import { LeftOutlined, PrinterOutlined, RightOutlined } from '@ant-design/icons';
import { Button, Card, Col, message, Row } from 'antd';
import { PageLoading } from '@ant-design/pro-layout';
import { useIntl } from 'umi';
import { print } from '../../services/device/index.js';
import styles from './index.less';

export const QueryPageFramework = (props) => {
  const intl = useIntl()

  const { type, dispatch, loading } = props;

  const printOnClick = () => {
    print().then((response) => {
      if (response.status == 200) {
        message.success(intl.formatMessage({ id: "pages.table.printSucceed" }));
      } else {
        message.error(intl.formatMessage({ id: "pages.table.printFailed" }));
      }
    });
  };

  const previousOnClick = () => {
    dispatch({
      type: 'device/previousPage',
      payload: {
        type: type,
      }
    });
  };

  const nextOnClick = () => {
    dispatch({
      type: 'device/nextPage',
      payload: {
        type: type,
      }
    });
  };


  if (loading) {
    return <PageLoading />;
  }

  return (
    <Card
      bordered={false}
      className={styles.card}
    >
      <Row justify="space-between" align="middle" gutter={16}>
        <Col span={1}>
          <Button
            shape="circle"
            icon={<LeftOutlined />}
            onClick={previousOnClick}
            style={{ boxShadow: '0 0.1rem 0.15rem 0 rgba(0, 0, 0, 0.35)' }}
          />
        </Col>
        <Col span={14} style={{ padding: 6, textAlign: 'center' }}>
          <div>
            {props.children}
          </div>
          <div style={{ padding: 25, textAlign: 'center' }}>
            <Button shape="round" icon={<PrinterOutlined />} onClick={printOnClick}
              style={{ boxShadow: '0 0.1rem 0.15rem 0 rgba(0, 0, 0, 0.35)' }}>
              {intl.formatMessage({ id: "pages.table.print" })}
            </Button>
          </div>
        </Col>
        <Col span={1}>
          <Button
            shape="circle"
            icon={<RightOutlined />}
            onClick={nextOnClick}
            style={{ boxShadow: '0 0.1rem 0.15rem 0 rgba(0, 0, 0, 0.35)' }} />
        </Col>
      </Row>
    </Card>
  );
};