-- 受外键约束反向 DROP 表
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS violation;
DROP TABLE IF EXISTS violation_kind;
DROP TABLE IF EXISTS bus;
DROP TABLE IF EXISTS line_staff;
DROP TABLE IF EXISTS line_station;
DROP TABLE IF EXISTS line;
DROP TABLE IF EXISTS fleet;
DROP TABLE IF EXISTS staff;
DROP TABLE IF EXISTS station;
-- 站点表
DROP TABLE IF EXISTS station;
CREATE TABLE station(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    name VARCHAR(50) COMMENT '站点名称',
    open_time DATETIME COMMENT '开通时间',
    close_time DATETIME COMMENT '关闭时间',
    PRIMARY KEY (id)
) COMMENT = '站点';
-- 员工表
DROP TABLE IF EXISTS staff;
CREATE TABLE staff(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    number VARCHAR(15) COMMENT '员工编号',
    name VARCHAR(20) COMMENT '姓名',
    sex VARCHAR(3) COMMENT '性别',
    native_place VARCHAR(255) COMMENT '籍贯',
    id_number VARCHAR(18) COMMENT '身份证号',
    phone VARCHAR(16) COMMENT '手机号',
    position VARCHAR(90) COMMENT '职位',
    wages DECIMAL(24, 2) COMMENT '工资',
    office VARCHAR(255) COMMENT '办公室',
    entry_time DATETIME COMMENT '入职时间',
    departure_time DATETIME COMMENT '离职时间',
    PRIMARY KEY (id)
) COMMENT = '员工';
-- 车队表
DROP TABLE IF EXISTS fleet;
CREATE TABLE fleet(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    name VARCHAR(255) NOT NULL COMMENT '名称',
    captain bigint unsigned COMMENT '队长',
    establish DATETIME COMMENT '成立时间',
    PRIMARY KEY (id),
    FOREIGN KEY (captain) REFERENCES staff(id),
    FOREIGN KEY (exterprise) REFERENCES exterprise(id)
) COMMENT = '车队';
-- 线路表
DROP TABLE IF EXISTS line;
CREATE TABLE line(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    name VARCHAR(255) COMMENT '名称',
    fleet bigint unsigned COMMENT '车队',
    open_time DATETIME COMMENT '开通时间',
    close_time DATETIME COMMENT '关闭时间',
    captain bigint unsigned COMMENT '路队长',
    PRIMARY KEY (id),
    FOREIGN KEY (fleet) REFERENCES fleet(id),
    FOREIGN KEY (captain) REFERENCES staff(id)
) COMMENT = '线路';
-- 路线-站点表
DROP TABLE IF EXISTS line_station;
CREATE TABLE line_station(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    line bigint unsigned COMMENT '路线',
    station bigint unsigned COMMENT '站点',
    next bigint unsigned COMMENT '下一站',
    PRIMARY KEY (id),
    FOREIGN KEY (line) REFERENCES line(id),
    FOREIGN KEY (station) REFERENCES station(id),
    FOREIGN KEY (next) REFERENCES station(id)
) COMMENT = '路线-站点';
-- 线路-员工表
DROP TABLE IF EXISTS line_staff;
CREATE TABLE line_staff(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    line bigint unsigned COMMENT '线路',
    staff bigint unsigned COMMENT '员工',
    PRIMARY KEY (id),
    FOREIGN KEY (line) REFERENCES line(id),
    FOREIGN KEY (staff) REFERENCES staff(id)
) COMMENT = '线路-员工';
-- 车辆表
DROP TABLE IF EXISTS bus;
CREATE TABLE bus(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    model VARCHAR(90) COMMENT '车辆型号',
    people INT COMMENT '核载人数',
    color VARCHAR(10) COMMENT '颜色',
    plate VARCHAR(10) COMMENT '车牌',
    line bigint unsigned COMMENT '线路',
    buy_time DATETIME COMMENT '购入时间',
    buy_by bigint unsigned COMMENT '购入人',
    PRIMARY KEY (id),
    FOREIGN KEY (line) REFERENCES line(id),
    FOREIGN KEY (buy_by) REFERENCES staff(id)
) COMMENT = '车辆';
-- 违章类型表
DROP TABLE IF EXISTS violation_kind;
CREATE TABLE violation_kind(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    name VARCHAR(255) COMMENT '名称',
    PRIMARY KEY (id)
) COMMENT = '违章类型';
-- 违章记录表
DROP TABLE IF EXISTS violation;
CREATE TABLE violation(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    liable_person bigint unsigned COMMENT '责任人',
    bus bigint unsigned COMMENT '车辆',
    station bigint unsigned COMMENT '违章站点',
    violation_time DATETIME COMMENT '违章时间',
    kind bigint unsigned COMMENT '违章类型',
    input_by bigint unsigned COMMENT '录入人',
    input_time DATETIME COMMENT '录入时间',
    PRIMARY KEY (id),
    FOREIGN KEY (liable_person) REFERENCES staff(id),
    FOREIGN KEY (bus) REFERENCES bus(id),
    FOREIGN KEY (station) REFERENCES station(id),
    FOREIGN KEY (kind) REFERENCES violation_kind(id),
    FOREIGN KEY (input_by) REFERENCES staff(id)
) COMMENT = '违章记录';
-- 登陆账号表
DROP TABLE IF EXISTS account;
CREATE TABLE account(
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    number bigint unsigned COMMENT '员工编号',
    password VARCHAR(255) COMMENT '密码',
    PRIMARY KEY (id),
    FOREIGN KEY (number) REFERENCES staff(id)
) COMMENT = '登陆账号';
-- 司机基本信息表
DROP VIEW driver;
CREATE VIEW driver AS
SELECT s.id AS id,
    s.number AS number,
    s.name AS name,
    s.sex AS sex,
    s.native_place AS native_place,
    s.id_number AS id_number,
    s.phone AS phone,
    s.position AS position,
    s.wages AS wages,
    s.office AS office,
    s.entry_time AS entry_time,
    s.departure_time AS departure_time,
    l.id AS line_id,
    l.name AS line,
    f.id AS fleet_id,
    f.name AS fleet
FROM staff AS s
    JOIN line_staff AS ls ON ls.staff = s.id
    JOIN line AS l ON l.id = ls.line
    JOIN fleet AS f ON f.id = l.fleet;
-- 司机违章表
DROP VIEW driver_violation;
CREATE VIEW driver_violation AS
SELECT s1.id AS id,
    s1.number AS number,
    s1.name AS name,
    f.name AS fleet,
    l.name AS line,
    st.name AS station,
    b.plate AS bus_plate,
    vk.name AS violation,
    v.violation_time AS violation_time,
    s2.name AS input_by,
    v.input_time AS input_time
FROM violation AS v
    JOIN staff AS s1 ON s1.id = v.liable_person
    JOIN line_staff AS ls ON ls.staff = s1.id
    JOIN line AS l ON l.id = ls.line
    JOIN fleet AS f ON f.id = l.fleet
    JOIN station AS st ON st.id = v.station
    JOIN bus AS b ON b.id = v.bus
    JOIN violation_kind AS vk ON vk.id = v.kind
    JOIN staff AS s2 ON s2.id = v.input_by;