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

CREATE TABLE tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    refresh_expired_at INT,
    user_id INT,
    access_token VARCHAR(255),
    refresh_token VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
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


CREATE TABLE food_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,  -- 음식 이름으로 이미지 참조
    image VARCHAR(255) DEFAULT 'food_default.png',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
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

-- 음식 선택했을 때 저장해야 된다. user_id = 1, time_id = 1, type_id = 1, scenario_id = 1, name = '김치찌개'
CREATE TABLE foods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL,
    food_image_id INT,
    time_id INT,
    type_id INT,
    scenario_id INT,
    theme_id INT,
    flavor_id INT,
    FOREIGN KEY (time_id) REFERENCES times(id),
    FOREIGN KEY (type_id) REFERENCES types(id),
    FOREIGN KEY (scenario_id) REFERENCES scenarios(id),
    FOREIGN KEY (theme_id) REFERENCES themes(id),
    FOREIGN KEY (flavor_id) REFERENCES flavors(id),
    FOREIGN KEY (food_image_id) REFERENCES food_images(id)
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



-- meta table에 types, scenarios, times, themes, flavors 테이블을 저장하는 sql 문 만들어줘
INSERT INTO meta_tables (table_name, table_description) VALUES ('types', '종류별'),('scenarios', '상황별'),('times', '시간별'),('themes', '기분/테마별'),('flavors', '맛별');

-- times 테이블에 아침, 점심, 저녁, 브런치, 간식, 야식 순으로 저장하는 sql 문 만들어줘
INSERT INTO times (name, description,image) VALUES ('전체', '전체','category_default.png'),('아침', '아침','times/breakfast.png'), ('점심', '점심','times/lunch.png'), ('저녁', '저녁','times/dinner.png'), ('브런치', '브런치','times/brunch.png'), ('간식', '간식','times/snack.png'), ('야식', '야식','times/late night snack.png');

-- types 테이블에 한식, 중식, 일식, 양식, 분식,베트남 음식, 인도 음식, 패스트 푸드, 디저트, 퓨전 요리 순으로 저장하는 sql 문 만들어줘
INSERT INTO types (name, description,image) VALUES ('전체', '전체','category_default.png'),('한식', '한식','types/korean food.png'), ('중식', '중식','types/chinese food.png'), ('일식', '일식','types/japanese food.png'), ('양식', '양식','types/western food.png'), ('분식', '분식','types/korean street food.png'), ('베트남 음식', '베트남 음식','types/vietnamese food.png'), ('인도 음식', '인도 음식','types/indian food.png'), ('패스트 푸드', '패스트 푸드','types/fast food.png'), ('디저트', '디저트','types/dessert.png'), ('퓨전 요리', '퓨전 요리','types/fusion cuisine.png');

-- scenarios 테이블에 연인, 혼반, 가족, 다이어트, 회식, 친구 순으로 저장하는 sql 문 만들어줘
INSERT INTO scenarios (name, description,image) VALUES ('전체', '전체','category_default.png'),('연인', '연인','scenarios/couple.png'), ('혼밥', '혼밥','scenarios/eating alone.png'), ('가족', '가족','scenarios/family.png'), ('다이어트', '다이어트','scenarios/diet.png'), ('회식', '회식','scenarios/company dinner.png'), ('친구', '친구','scenarios/friend.png');

-- flavors 테이블에 매운맛, 감칠맛, 고소한맛, 단맛, 짠맛, 싱거운맛 순으로 저장하는 sql 문 만들어줘
INSERT INTO flavors (name, description,image) VALUES ('전체', '전체','category_default.png'),('매운맛', '매운맛','flavors/spicy taste.png'), ('감칠맛', '감칠맛','flavors/umami.png'), ('고소한맛', '고소한맛','flavors/nutty taste.png'), ('단맛', '단맛','flavors/sweet taste.png'), ('짠맛', '짠맛','flavors/salty taste.png'), ('싱거운맛', '싱거운맛','flavors/bland taste.png');

-- themes 테이블에 스트레스 해소, 피로 회복, 기분 전환, 제철 음식, 영양식, 특별한 날 순으로 저장하는 sql 문 만들어줘
INSERT INTO themes (name, description,image) VALUES ('전체', '전체','category_default.png'),('스트레스 해소', '스트레스 해소','themes/stress.png'), ('피로 회복', '피로 회복','themes/fatigue recovery.png'), ('특별한 날', '특별한 날','themes/mood refresh.png'), ('제철 음식', '제철 음식','themes/seasonal food.png');

INSERT INTO users (email, password, name,birth, sex, provider) VALUES ('test@jokertrickster.com', 'asdasd123', '푸드픽맨','1990-01-01', 'male', 'test');
