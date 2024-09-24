# Food Recommendation Service

## 회의 일지 및 프로젝트 문서 링크
### [회의 일지 열기](https://github.com/JokerTrickster/food-recommendation/wiki/%ED%9A%8C%EC%9D%98-%EC%9D%BC%EC%A7%80)
### [문서 링크 열기](https://github.com/JokerTrickster/food-recommendation/wiki/%EB%AC%B8%EC%84%9C-%EB%A7%81%ED%81%AC)

## Introduction
This project is a food recommendation service designed to help users find food options based on their preferences. The service uses advanced algorithms to suggest meals, making dining choices easier and more personalized.

## Requirements
To run the server, you will need the following libraries and tools:

- [Echo v4](https://github.com/labstack/echo) - High performance, minimalist Go web framework
- [Air v1.52](https://github.com/cosmtrek/air) - Live reload for Go apps
- [Docker](https://www.docker.com/) - Platform to develop, ship, and run applications
- [echo-swagger](https://github.com/swaggo/echo-swagger) - Swagger integration with Echo for API documentation

## Architecture
This project follows the principles of Clean Architecture. This architectural pattern emphasizes the separation of concerns, making the codebase more modular, testable, and maintainable. The core idea is to keep the business logic independent of frameworks, databases, and external agencies.

<p align="center">
    <img width="990" alt="스크린샷 2022-12-22 오후 7 46 07" src="https://user-images.githubusercontent.com/35329247/209118510-3153c568-0d17-43de-a778-210dd53002c5.png">
</p>

## DB Schema
![음식 추천 db](https://github.com/user-attachments/assets/333fa72e-8379-461e-a111-fcfaea260cbc)