## Automate the Sainsbury's order

### Building

`go build -o food.exe`

### Usage

Edit `data/` to contain correct recipes, then build.

`food auth login`  - authenticate
`food auth status` - check login succeeded
`food plan`        - plan order
`food validate`    - validate plan is all available
`food order`       - order validated plan

### Rough plan
- Store recipes in yaml (recipe: ingredients)
- Store product catalogue in yaml (for each recipe, store the Sainsbury's UID we want to buy)
- Give user option to select items they don't want to order
- Create plan for what actions we will take on Sainsbury's
- Get user to log in with playwright, and get authenticated session + delivery ID
- Using unofficial API, determine whether plan is actionable
- If not, ignore the bits that aren't actionable until we have a new plan, suggest this to user for validation
- Then add items to order
