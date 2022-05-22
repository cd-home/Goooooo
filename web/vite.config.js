import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import dotenv from 'dotenv'
import fs from 'fs'
import path from 'path'

// https://vitejs.dev/config/
export default ({command, mode}) => {
    const config = `./src/config`
    const envfs = [`${config}/.env`, `${config}/.env.${mode}`]
    for (const envf of envfs) {
        const envConfig = dotenv.parse(fs.readFileSync(envf))
        process.env = {...process.env, ...envConfig}
    }
    console.log(process.env.VITE_APP_ADDR)
    let server = {
        host: '127.0.0.1',
        port: 3000,
        open: false,
        https: false,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8080',
                changeOrigin: true,
                rewrite: path => path.replace('/^\/api/', '')
            }
        },
    }
    return defineConfig({
        plugins: [vue()],
        server,
        define: {
            "process.env": {...process.env}
        },
        resolve: {
            alias: {
                '@': path.resolve(__dirname, './src/')
            }
        }
    })
}
