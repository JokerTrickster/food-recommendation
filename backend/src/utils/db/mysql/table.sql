CREATE TABLE tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    refresh_expired_at INT,
    user_id INT,
    access_token VARCHAR(255),
    refresh_token VARCHAR(255)
);
-- email, 비밀번호, 생년월일, 성별 
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    email VARCHAR(255),
    name VARCHAR(255),
    password VARCHAR(255),
    birth DATE,
    sex varchar(50),
    provider VARCHAR(50)
);

-- 알레르기 정보 저장
create table allergies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) UNIQUE -- 알레르기 이름 (예: 'Peanut', 'Gluten', 'Lactose' 등)
    description VARCHAR(255)
);

-- 유저 알레르기 정보 저장
CREATE TABLE user_allergies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    user_id INT,
    allergy_id INT,
    UNIQUE KEY unique_user_allergy (user_id, allergy_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (allergy_id) REFERENCES allergies(id) ON DELETE CASCADE
);

-- 음식 선택했을 때 저장해야 된다. user_id = 1, time_id = 1, type_id = 1, scenario_id = 1, name = '김치찌개'
CREATE TABLE foods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) default 'food_default.png',
    time_id INT,
    type_id INT,
    scenario_id INT,
    theme_id INT,
    flavor_id INT
);

-- 유저에게 추천된 음식을 저장해야 된다.
CREATE TABLE food_histories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    food_id INT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (food_id) REFERENCES foods(id)
);
  
-- 메타 데이터 테이블
  CREATE TABLE meta_tables (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    table_name VARCHAR(255) NOT NULL UNIQUE,
    table_description VARCHAR(255)
);


-- 아침, 점심, 저녁, 브런치, 간식, 야식
CREATE TABLE times (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    image varchar(255) default 'category_default.png',
    description VARCHAR(255)
);

-- 한식, 중식, 일식, 양식, 분식,베트남 음식, 인도 음식, 패스트 푸드, 디저트, 퓨전 요리
CREATE TABLE types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    image varchar(255) default 'category_default.png',
    description VARCHAR(255)
);

-- 연인, 혼반, 가족, 다이어트, 회식, 친구
CREATE TABLE scenarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    image varchar(255) default 'category_default.png',
    description VARCHAR(255)
);

-- 매운맛, 감칠맛, 고소한맛, 단맛, 짠맛, 싱거운맛 
CREATE TABLE flavors(
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    image varchar(255) default 'category_default.png',
    description VARCHAR(255)    
);

-- 스트레스 해소, 피로 회복, 기분 전환, 제철 음식, 영양식, 특별한 날
CREATE TABLE themes(
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    image varchar(255) default 'category_default.png',
    description VARCHAR(255)    
);

CREATE TABLE user_auths (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    auth_code VARCHAR(255),
    email VARCHAR(255),
    type VARCHAR(100)
);

-- meta table에 types, scenarios, times, themes, flavors 테이블을 저장하는 sql 문 만들어줘
INSERT INTO meta_tables (table_name, table_description) VALUES ('types', '종류별'),('scenarios', '상황별'),('times', '시간별'),('themes', '기분/테마별'),('flavors', '맛별');

-- times 테이블에 아침, 점심, 저녁, 브런치, 간식, 야식 순으로 저장하는 sql 문 만들어줘
INSERT INTO times (name, description,image) VALUES ('아침', '아침','category_default.png'), ('점심', '점심','category_default.png'), ('저녁', '저녁','category_default.png'), ('브런치', '브런치','category_default.png'), ('간식', '간식','category_default.png'), ('야식', '야식','category_default.png');

-- types 테이블에 한식, 중식, 일식, 양식, 분식,베트남 음식, 인도 음식, 패스트 푸드, 디저트, 퓨전 요리 순으로 저장하는 sql 문 만들어줘
INSERT INTO types (name, description,image) VALUES ('한식', '한식','category_default.png'), ('중식', '중식','category_default.png'), ('일식', '일식','category_default.png'), ('양식', '양식','category_default.png'), ('분식', '분식','category_default.png'), ('베트남 음식', '베트남 음식','category_default.png'), ('인도 음식', '인도 음식','category_default.png'), ('패스트 푸드', '패스트 푸드','category_default.png'), ('디저트', '디저트','category_default.png'), ('퓨전 요리', '퓨전 요리','category_default.png');

-- scenarios 테이블에 연인, 혼반, 가족, 다이어트, 회식, 친구 순으로 저장하는 sql 문 만들어줘
INSERT INTO scenarios (name, description,image) VALUES ('연인', '연인','category_default.png'), ('혼반', '혼반','category_default.png'), ('가족', '가족','category_default.png'), ('다이어트', '다이어트','category_default.png'), ('회식', '회식','category_default.png'), ('친구', '친구','category_default.png');

-- flavors 테이블에 매운맛, 감칠맛, 고소한맛, 단맛, 짠맛, 싱거운맛 순으로 저장하는 sql 문 만들어줘
INSERT INTO flavors (name, description,image) VALUES ('매운맛', '매운맛','category_default.png'), ('감칠맛', '감칠맛','category_default.png'), ('고소한맛', '고소한맛','category_default.png'), ('단맛', '단맛','category_default.png'), ('짠맛', '짠맛','category_default.png'), ('싱거운맛', '싱거운맛','category_default.png');

-- themes 테이블에 스트레스 해소, 피로 회복, 기분 전환, 제철 음식, 영양식, 특별한 날 순으로 저장하는 sql 문 만들어줘
INSERT INTO themes (name, description,image) VALUES ('스트레스 해소', '스트레스 해소','category_default.png'), ('피로 회복', '피로 회복','category_default.png'), ('기분 전환', '기분 전환','category_default.png'), ('제철 음식', '제철 음식','category_default.png'), ('영양식', '영양식','category_default.png'), ('특별한 날', '특별한 날','category_default.png');

-- 알레르기 정보 저장
Insert INTO allergies (name, description) VALUES ('기타', '기타'), ('계란', '계란'), ('우유', '우유'), ('메밀', '메밀'), ('땅콩', '땅콩'), ('대두', '대두'), ('밀', '밀');

INSERT INTO users (email, password, name,birth, sex, provider) VALUES ('test@jokertrickster.com', 'asdasd123', '푸드픽맨','1990-01-01', 'male', 'test');
