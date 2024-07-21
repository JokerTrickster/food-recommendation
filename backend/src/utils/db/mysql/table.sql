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
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    birth DATE,
    sex varchar(50),
    provider VARCHAR(50)
);
-- 음식 선택했을 때 저장해야 된다. user_id = 1, time_id = 1, type_id = 1, scenario_id = 1, name = '김치찌개'
CREATE TABLE foods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL,
    time_id INT,
    type_id INT,
    scenario_id INT
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


-- 전체, 조식, 중식, 석식, 간식, 야식
CREATE TABLE times (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255)
);

-- 전체, 한식, 중식, 양식, 일식, 분식, 간식
CREATE TABLE types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255)
);

-- 전체, 혼밥, 친구, 가족, 회식 등
CREATE TABLE scenarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
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

-- meta table에 times, types, scenarios 테이블을 저장하는 sql 문 만들어줘
INSERT INTO meta_tables (table_name, table_description) VALUES ('times', '시간대'), ('types', '음식 종류'), ('scenarios', '상황');

-- times 테이블에 전체, 조식, 중식, 석식, 간식, 야식 순으로 저장하는 sql 문 만들어줘
INSERT INTO times (name, description) VALUES ('전체', '전체'), ('조식', '아침'), ('중식', '점심'), ('석식', '저녁'), ('간식', '간식'), ('야식', '야식');

-- types 테이블에 전체, 한식, 중식, 양식, 일식, 분식 순으로 저장하는 sql 문 만들어줘
INSERT INTO types (name, description) VALUES ('전체', '전체'), ('한식', '한식'), ('중식', '중식'), ('양식', '양식'), ('일식', '일식'), ('분식', '분식');

-- scenarios 테이블에 전체, 혼밥, 친구, 가족, 회식 순으로 저장하는 sql 문 만들어줘
INSERT INTO scenarios (name, description) VALUES ('전체', '전체'), ('혼밥', '혼밥'), ('친구', '친구'), ('연인','연인') ('가족', '가족'), ('회식', '회식');