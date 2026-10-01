## Automate the Sainsbury's order

### Rough plan
- Store recipes in yaml (recipe: ingredients)
- Store product catalogue in yaml (for each recipe, store the Sainsbury's UID we want to buy)
- Give user option to select items they don't want to order
- Create plan for what actions we will take on Sainsbury's
- Get user to log in with playwright, and get authenticated session + delivery ID
- Using unofficial API, determine whether plan is actionable
- If not, ignore the bits that aren't actionable until we have a new plan, suggest this to user for validation
- Then add items to order
