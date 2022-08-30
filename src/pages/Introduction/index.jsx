import { useState } from 'react';
import styles from './index.less'
import { Image, Col, Row, Card } from 'antd';
import { MailFilled, EnvironmentFilled } from '@ant-design/icons';

export default () => {
  return (
    <Col>
      <Row>
        <Card className={styles.headFirst} />
        <Row className={styles.headSecond} >
          <Col span={2} />
          <Col span={2}>
            <div className={styles.headLogo}>
              <Image
                preview={false}
                src="homeLogo.png"
                height={"102%"}
                width={"102%"}
              />
            </div>
          </Col>
          <Col span={1} />
          <Col span={8}>
            <Row className={styles.headTitle}>保定大恒电气科技有限公司</Row>
            <Row className={styles.headTitleUs}>Baoding Daheng Electric Technology Co. ,Ltd</Row>
          </Col>
        </Row>
      </Row>
      <Row>
        <Col className={styles.body}>
          <div className={styles.bodySubscriptBox}>
            <div className={styles.bodySubscriptBorder} />
            <div className={styles.bodySubscript} />
          </div>
          <Image
            className={styles.bodyBackgroundImage}
            preview={false}
            placeholder={true}
            src="body.jpg"
            height={"100%"}
            width={"100%"}
          />
          <div className={styles.bodyBanner}>
            <div className={styles.bodyBannerLeft}>
              <div className={styles.bodyBannerLeftBackground}>
                <div className={styles.bodyBannerLeftText}>CORPORATE PURSUITS</div>
              </div>
            </div>

            <div className={styles.bodyBannerMiddle}>
              <div className={styles.bodyBannerMiddleBackground}>
                <Row className={styles.bodyBannerMiddleIcon}>
                  <div className={styles.bodyBannerMiddleIconItem} />
                  <div className={styles.bodyBannerMiddleIconItem} />
                  <div className={styles.bodyBannerMiddleIconItem} />
                </Row>
              </div>
            </div>

            <div className={styles.bodyBannerRightBox}>
              <div className={styles.bodyBannerRight} />
              <div className={styles.bodyBannerRightBackground}>
                <div className={styles.bodyBannerRightText}>
                  大恒电气的追求目标：追求客户满意，创优质品牌，铸一流企业形象
                  <br />
                  DAHENG ELECTRIC'S PURSUIT GOAL: THE PURSUIT OF CUSTOMER SATISFACTION,
                  <br />
                  CREATE HIGH-QUALITY BRAND, CAST FIRST-CLASS CORPORATE IMAGE.
                  <br />
                </div>
              </div>
            </div>
          </div>

          <div className={styles.bodyContent}>
            <Row>
              <Col span={7}>
                <Row>
                  <Row className={styles.bodyContentIntroductionLeft}>
                    公司
                  </Row>
                  <Row className={styles.bodyContentIntroductionRight}>
                    简介
                  </Row>
                </Row>
                <Row className={styles.bodyContentIntroductionUS}>
                  COMPANY PROFILE
                </Row>
                <Row className={styles.bodyContentIntroductionContent}>
                  保定大恒电气科技有限公司，现位于保定市高新产业技术开发区内，我公司是一家自主研发，自主生产，多元化销售为一体的高新技术产业。
                </Row>
              </Col>
              <Col span={1} />
              <Col span={8}>
                <Image
                  className={styles.bodyContentIntroductionImage}
                  preview={false}
                  placeholder={true}
                  src="synopsis.jpg"
                  height={"66%"}
                  width={"100%"}
                />
              </Col>
              <Col span={1} />
              <Col span={6} className={styles.bodyContentIntroductionDeviceList}>
                <Col onClick={() => {
                  shell.open("http://www.dahengdianqi.com")
                }}>
                  <Row className={styles.bodyContentIntroductionDeviceListItem}>
                    <div className={styles.bodyContentIntroductionDeviceListIcon} />
                    油化检测设备
                  </Row>
                  <Row className={styles.bodyContentIntroductionDeviceListItem}>
                    <div className={styles.bodyContentIntroductionDeviceListIcon} />
                    器皿清洗设备
                  </Row>
                  <Row className={styles.bodyContentIntroductionDeviceListItem}>
                    <div className={styles.bodyContentIntroductionDeviceListIcon} />
                    高压检测设备
                  </Row>
                  <Row>
                    <Col>
                      <div className={styles.bodyContentIntroductionDeviceListIcon} />
                    </Col>
                    <Col className={styles.bodyContentIntroductionDeviceListItem}>
                      <div>
                        更多设备
                      </div>
                      <p className={styles.bodyContentIntroductionDeviceListLink} >
                        www.dahengdianqi.com
                      </p>
                    </Col>
                  </Row>
                </Col>
              </Col>
              <Col span={1}>
                <Image
                  className={styles.bodyContentIntroductionImage}
                  preview={false}
                  placeholder={true}
                  src="QRcode.jpg"
                  height={"10vh"}
                  width={"10vh"}
                />
                <div className={styles.bodyContentIntroductionContact}>联系人：于经理</div>
                <div className={styles.bodyContentIntroductionPhone}>销售热线：181-3238-4396</div>
              </Col>
            </Row>
          </div>
        </Col>
      </Row >
      <Row>
        <Row className={styles.footFirst}>
          <Col style={{ color: "#FFFFFF", textAlign: "right", textJustify: "initial" }} span={11}>
            <MailFilled className={styles.footIcon} />
            bddaheng@163.com
          </Col>
          <Col span={2} />
          <Col style={{ color: "#FFFFFF", textAlign: "left" }} span={11}>
            <EnvironmentFilled className={styles.footIcon} />
            保定市高新区锦绣街677号火炬产业园
          </Col>
        </Row>
        <Row className={styles.footSecond} />
      </Row>
    </Col >
  );
};