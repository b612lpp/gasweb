package logic

const totalValue = 1500

type ToFillTankType struct{ ToFullLiters, ToFullCost int }

func DoMath(v, c int) (ToFillTank ToFillTankType, err error) {
	tl := ((85 - v) * totalValue) / 100 //считаем литры до полного бака
	tc := tl * c                        //умножаем литры на цену
	ToFillTank = ToFillTankType{ToFullLiters: tl, ToFullCost: tc}
	return ToFillTank, err
}
