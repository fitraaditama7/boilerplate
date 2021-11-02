INSERT INTO tm_user 
(id, username, email, password, first_name, last_name, phone_number, role_id, created_at, created_by)
VALUES
('8e0f4c66-9782-4e09-a311-6e662d11222d', 'user1', 'user1@mail.com', '$2a$04$zOuUaVfOgRQ6UmGRqHnJhu7kxtS4LztpIaDgSmuYWOEUTV2p4eLCK', 'user1', 'user1', '081234569', 'admin', NOW(), 'fitra'),
('efd18409-1e9c-4810-bb21-1fc1c1b37ba8', 'user2', 'user2@mail.com', '$2a$04$zOuUaVfOgRQ6UmGRqHnJhu7kxtS4LztpIaDgSmuYWOEUTV2p4eLCK', 'user2', 'user2', '081234568', 'admin', NOW(), 'fitra'),
('86f0e528-dfb9-440c-a4fd-15572b9fff32', 'user3', 'user3@mail.com', '$2a$04$zOuUaVfOgRQ6UmGRqHnJhu7kxtS4LztpIaDgSmuYWOEUTV2p4eLCK', 'user3', 'user3', '081234567', 'admin', NOW(), 'fitra'),
('07a2bf91-c9ce-438a-a11f-a7ac54793c98', 'user4', 'user4@mail.com', '$2a$04$zOuUaVfOgRQ6UmGRqHnJhu7kxtS4LztpIaDgSmuYWOEUTV2p4eLCK', 'user4', 'user4', '081234566', 'admin', NOW(), 'fitra'),
('1c4f0a35-fde3-4b66-92d4-ed9f10312468', 'user5', 'user5@mail.com', '$2a$04$zOuUaVfOgRQ6UmGRqHnJhu7kxtS4LztpIaDgSmuYWOEUTV2p4eLCK', 'user5', 'user5', '081234565', 'admin', NOW(), 'fitra')
ON DUPLICATE KEY UPDATE
  updated_at = NOW();
