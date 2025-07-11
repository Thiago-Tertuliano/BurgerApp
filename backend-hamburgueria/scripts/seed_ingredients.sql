-- Script para inserir ingredientes de exemplo
-- Execute este script no PostgreSQL para popular a tabela ingredients

-- Limpar dados existentes (opcional)
-- DELETE FROM ingredients;

-- Inserir pães
INSERT INTO ingredients (name, price, category, is_available) VALUES
('Pão Brioche', 2.00, 'pao', true),
('Pão Australiano', 2.50, 'pao', true),
('Pão Tradicional', 1.50, 'pao', true),
('Pão de Hambúrguer', 2.00, 'pao', true);

-- Inserir carnes
INSERT INTO ingredients (name, price, category, is_available) VALUES
('Hambúrguer Bovino 120g', 6.00, 'carne', true),
('Hambúrguer de Frango', 5.50, 'carne', true),
('Hambúrguer Vegano', 7.00, 'carne', true),
('Hambúrguer de Peixe', 6.50, 'carne', true);

-- Inserir queijos
INSERT INTO ingredients (name, price, category, is_available) VALUES
('Queijo Cheddar', 2.00, 'queijo', true),
('Queijo Prato', 2.00, 'queijo', true),
('Queijo Muçarela', 2.00, 'queijo', true),
('Queijo Gouda', 2.50, 'queijo', true);

-- Inserir vegetais/acompanhamentos
INSERT INTO ingredients (name, price, category, is_available) VALUES
('Alface', 1.00, 'vegetal', true),
('Tomate', 1.00, 'vegetal', true),
('Cebola', 1.00, 'vegetal', true),
('Bacon', 2.50, 'acompanhamento', true),
('Cebola Crispy', 1.50, 'acompanhamento', true),
('Picles', 1.00, 'acompanhamento', true);

-- Inserir molhos
INSERT INTO ingredients (name, price, category, is_available) VALUES
('Maionese da Casa', 1.00, 'molho', true),
('Barbecue', 1.00, 'molho', true),
('Mostarda e Mel', 1.00, 'molho', true),
('Ketchup', 1.00, 'molho', true);

-- Verificar os dados inseridos
SELECT * FROM ingredients ORDER BY category, name; 