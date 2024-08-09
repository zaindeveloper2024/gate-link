import express from 'express'

const app1 = express()
const app2 = express()
const PORT1 = 3000
const PORT2 = 3001

app1.get('/api/users', (req, res) => {
    res.json({
        message: `Hello from Port ${PORT1}`
    })
})

app1.get('/api/users/error', (req, res) => {
    throw new Error('Error from Port 3000')
})

app2.get('/api/products', (req, res) => {
    res.json({
        message: `Hello from Port ${PORT2}`
    })
})

app1.listen(PORT1, () => {
    console.log(`Server1 is running on port ${PORT1}`)
})

app2.listen(PORT2, () => {
    console.log(`Server2 is running on port ${PORT2}`)
})