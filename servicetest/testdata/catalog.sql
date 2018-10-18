INSERT INTO service (id, service_id, name, description, tags, requires) VALUES (1, '19388425-2619-424e-820c-cd575465c657', 'test-service-1', 'a test service', '["one", "two", "three"]', '{"one":"two"}');
INSERT INTO service (id, service_id, name, description, tags, requires) VALUES (2, '19388425-2619-424e-820c-cd575465c658', 'test-service-2', 'a test service', '["one", "two", "three"]', '{"one":"two"}');
INSERT INTO service (id, service_id, name, description, tags, requires) VALUES (3, '19388425-2619-424e-820c-cd575465c659', 'test-service-3', 'a test service', '["one", "two", "three"]', '{"one":"two"}');

INSERT INTO service_plan (id, plan_id, name, service_id) VALUES (1, '19388425-2619-424e-820c-cd575465c650', 'test-service-plan-1', 1);
INSERT INTO service_plan (id, plan_id, name, service_id) VALUES (2, '19388425-2619-424e-820c-cd575465c651', 'test-service-plan-2', 1);
INSERT INTO service_plan (id, plan_id, name, service_id) VALUES (3, '19388425-2619-424e-820c-cd575465c652', 'test-service-plan-3', 2);
