#!/usr/bin/env node
// This tells our computer that
// this is a node script rather than bash script.
'use strict';

const amqp = require('amqplib');

const qName = 'testQ';
const mqAddr = process.env.MQADDR || 'localhost:5672';
const mqURL = `amqp://${mqAddr}`;

(async function() {
    try {
        console.log('connecting to %s', mqURL);
        let connection = await amqp.connect(mqURL);
        let channel = await connection.createChannel();
        // Durable queue writes messages to disk.
        // So even our MQ server dies,
        // the information is saved on disk and not lost.
        let qConf = await channel.assertQueue(qName, { durable: false });

        setInterval(() => {
            let data = {
                name: 'Zico Deng',
                age: 21
            };

            console.log('sending messages:', data);

            // Data needs to be serialized before handing it to MQ.
            channel.sendToQueue(qName, Buffer.from(JSON.stringify(data)));
        }, 1000);
    } catch (err) {
        console.log(err.stack);
    }
})();
