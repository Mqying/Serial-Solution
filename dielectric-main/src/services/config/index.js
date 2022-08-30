import { request } from 'umi'
const sleep = ms => new Promise(r => setTimeout(r, ms));

export async function getConfig() {
    let retry = 0;

    while (retry < 20) {
        try {
            const response = await request('/api/v1/config/get', {
                method: 'GET',
                prefix: 'http://127.0.0.1:9090'
            })
            return response
        } catch (err) {
            if (err.response === null) {
                retry++;
                await sleep(1000)
                continue
            }

            return {
                status: 500,
                config: {
                    dielectron: 0,
                    water: 0,
                    acid: 0,
                    flash: 0,
                    logo: 0
                }
            }
        }
    }

    return {
        status: 500,
        config: {
            dielectron: 0,
            water: 0,
            acid: 0,
            flash: 0,
            logo: 0
        }
    }
}