package custom_queries

import (
	"gorm.io/gorm"
	"log"
)

func CreateTriggerIfNotExists(db *gorm.DB) {
	query := `CREATE OR REPLACE FUNCTION update_budget_utilization()
    RETURNS TRIGGER AS $$
BEGIN
    -- Handle DELETE: Remove the transaction amount from utilization
    IF TG_OP = 'DELETE' THEN
        UPDATE budgets
        SET utilization = COALESCE(utilization - OLD.amount, 0),
        fraction = (COALESCE(utilization - OLD.amount, 0) / amount) * 100,
        balance = balance + OLD.amount
        WHERE user_id = OLD.user_id AND category_id = OLD.category_id;
        RETURN OLD;
    END IF;

    -- Handle UPDATE: Remove the old amount, then add the new amount
    IF TG_OP = 'UPDATE' THEN
        UPDATE budgets
        SET utilization = utilization - OLD.amount + NEW.amount,
        fraction = ((utilization - OLD.amount + NEW.amount) / amount) * 100,
        balance = balance - OLD.amount + NEW.amount
        WHERE user_id = OLD.user_id AND category_id = OLD.category_id;
        RETURN NEW;
    END IF;
    -- Handle INSERT: Add the new transaction amount to utilization
    UPDATE budgets
    SET utilization = utilization + NEW.amount,
    fraction = ((utilization + NEW.amount )/ amount) * 100,
    balance = balance - NEW.amount
    WHERE user_id = NEW.user_id AND category_id = NEW.category_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_budget_utilization_trigger
    AFTER INSERT OR UPDATE OR DELETE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_budget_utilization();
`
	result := db.Exec(query)
	if result.Error != nil {
		log.Fatalln("Could not create trigger " + result.Error.Error())

	}
	log.Println("Trigger created successfully")

}
