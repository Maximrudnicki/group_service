## Pardo: An Online Platform for Learning Vocabulary

Pardo is an online platform created to facilitate the learning of foreign vocabulary. Initially conceived as a start-up project in the summer of 2021 with students from my university, it was later chosen as my diploma project and began active development in December 2023. Pardo serves not only individuals learning new languages but also teachers and their students by allowing the creation and management of groups or joining existing ones.

### Microservices

Pardo is composed of several microservices:

1. [Broker](https://github.com/Maximrudnicki/broker-service-pardo)
2. [Vocabulary](https://github.com/Maximrudnicki/vocab-service-pardo)
3. [Auth](https://github.com/Maximrudnicki/auth-service-pardo)
4. [Group](https://github.com/Maximrudnicki/group_service)
5. [Mail](https://github.com/Maximrudnicki/mail-service-pardo)

### Front-End

The front-end code can be found [here](https://github.com/Maximrudnicki/vue_pardo). You can also use the platform online by registering [here](https://vue-pardo.onrender.com/).

### Python Microservices

Additionally, there are two microservices built using Python:

1. [Py-Vocabulary](https://github.com/Maximrudnicki/py-vocab-pardo)
2. [Py-Broker](https://github.com/Maximrudnicki/py-broker-pardo)

### Tech Stack

The following technologies were used to build this project:

- **Languages**: Golang, Python, JavaScript
- **Frameworks**: Gin, FastAPI, gRPC
- **Front-End Framework**: Vue.js
- **Databases**: MongoDB, PostgreSQL
- **ORMs**: GORM, SQLAlchemy, Alembic (migration tool for SQLAlchemy)
- **Message Broker**: RabbitMQ
- **Protocols**: Protocol Buffers, JSON
