CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_sender_id ON transactions(sender_id);
CREATE INDEX idx_transactions_receiver_id ON transactions(receiver_id);
CREATE INDEX idx_transactions_type ON transactions(type);