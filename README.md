# Food Recommendation Service

## 회의 일지 및 프로젝트 문서 링크
### [앱 스토어 링크](https://play.google.com/store/apps/details?id=com.food_pick&hl=ko)
### [회의 일지 열기](https://github.com/JokerTrickster/food-recommendation/wiki/%ED%9A%8C%EC%9D%98-%EC%9D%BC%EC%A7%80)
### [문서 링크 열기](https://github.com/JokerTrickster/food-recommendation/wiki/%EB%AC%B8%EC%84%9C-%EB%A7%81%ED%81%AC)

## Introduction

ai 를 사용하여 음식 궁합과 사용자 취향에 맞는 음식 추천 서비스

## App Screenshot

<div style="display: flex; flex-wrap: nowrap; gap: 10px;">
  <div style="text-align: center;">
    <img src="https://github.com/user-attachments/assets/356b68bd-6676-4f46-b67b-ae13341d4565" alt="메인화면" width="150"/>
    <p>메인화면</p>
  </div>
  <div style="text-align: center;">
    <img src="https://github.com/user-attachments/assets/c668a5aa-4f3b-4746-b19d-0bf42d6aa990" alt="데일리 추천 화면" width="150"/>
    <p>데일리 추천 화면</p>
  </div>
  <div style="text-align: center;">
    <img src="https://github.com/user-attachments/assets/5a86dc39-14af-44c1-88d0-a70d63e64eec" alt="음식 선택 화면" width="150"/>
    <p>음식 선택 화면</p>
  </div>
  <div style="text-align: center;">
    <img src="https://github.com/user-attachments/assets/6097df52-9aa9-4925-a9ac-1fb7f4f490ee" alt="음식 궁합 화면" width="150"/>
    <p>음식 궁합 화면</p>
  </div>
  <div style="text-align: center;">
    <img src="https://github.com/user-attachments/assets/df820499-7dc4-4b3a-a2af-059d966ce36b" alt="음식 랜덤 선택 화면" width="150"/>
    <p>음식 랜덤 선택 화면</p>
  </div>
</div>

## MVP Features
1. 인증 기능
2. 데일리 추천 기능
3. 음식 추천 기능
4. 음식 선택시 궁합 기능
5. 사용자 취향에 맞는 음식 추천 앱 푸시 기능


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
<p align="center">
   <img width="656" alt="Image" src="https://github.com/user-attachments/assets/21a101ab-4bd9-4bd1-a501-bde833bec0ba" />
</p>