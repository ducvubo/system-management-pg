-- name: CreateUserPatoProfile :execresult
INSERT INTO user_pato_profile (
     us_pt_name, us_pt_avatar, us_pt_phone,
    us_pt_gender, us_pt_address, us_pt_birthday, createdAt
) VALUES (
 ?, ?, ?, ?, ?, ?, NOW()
);
