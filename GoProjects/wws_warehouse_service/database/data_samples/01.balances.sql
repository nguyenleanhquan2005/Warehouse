INSERT INTO balances (user_id, currency, amount, version, created_at, updated_at, deleted_at)
VALUES
-- User A: regular casual player
('user_001', 'chips',    15000, 1, NOW(), NOW(), NULL),
('user_001', 'gold',      3000, 1, NOW(), NOW(), NULL),
('user_001', 'diamonds',   120, 1, NOW(), NOW(), NULL),

-- User B: high roller
('user_002', 'chips',   250000, 3, NOW(), NOW(), NULL),
('user_002', 'gold',     12000, 2, NOW(), NOW(), NULL),
('user_002', 'diamonds',   980, 1, NOW(), NOW(), NULL),
('user_002', 'tickets',     20, 1, NOW(), NOW(), NULL),

-- User C: new player
('user_003', 'chips',     5000, 1, NOW(), NOW(), NULL),
('user_003', 'gold',       500, 1, NOW(), NOW(), NULL),

-- User D: inactive player
('user_004', 'chips',     8000, 1, NOW() - INTERVAL 60 DAY, NOW() - INTERVAL 60 DAY, NULL),
('user_004', 'gold',      1000, 1, NOW() - INTERVAL 60 DAY, NOW() - INTERVAL 60 DAY, NULL),
('user_004', 'diamonds',    40, 1, NOW() - INTERVAL 60 DAY, NOW() - INTERVAL 60 DAY, NULL),

-- User E: recently deleted
('user_005', 'chips',    10000, 1, NOW() - INTERVAL 30 DAY, NOW() - INTERVAL 20 DAY, NOW() - INTERVAL 10 DAY),
('user_005', 'gold',      2500, 1, NOW() - INTERVAL 30 DAY, NOW() - INTERVAL 20 DAY, NOW() - INTERVAL 10 DAY);
