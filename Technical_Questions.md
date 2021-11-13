# Technical Questions

### 1.I used 4 hours for test(include add more feature).I improve my funcion test by add sort feature by performance (max,min)
### 2.Sort feature is help user. User can view a list of funds by sort max or min performance in day, week, month and year
```
	//add sort performance feature
	if s != "" {
		for _, v := range response.Data {
			sort.Slice(v.Value, func(i, j int) bool {
				if strings.ToUpper(s) == model.Min {
					return v.Value[i].Performance < v.Value[j].Performance
				}
				return v.Value[i].Performance > v.Value[j].Performance
			})
		}
	}
  ```
  How to use 
  ```
go run main.go -range=1W sort=max
  ```
  ![GitHub Logo](https://github.com/bankff/finnomenaTest/blob/main/src/exwithsort.png)


### 3. I track issue in production by logs  
### 4. First, I would improve the FINNOMENA APIs by add sort feature by nav_date and performance, last I would improve the APIs by add feature filter by range time, thailand_fund_code, range price (Nov)
