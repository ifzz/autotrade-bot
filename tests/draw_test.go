package tests

import (
	"math"
	"testing"

	"github.com/rwlist/autotrade-bot/formula"
	"github.com/stretchr/testify/assert"

	"github.com/rwlist/autotrade-bot/draw"
)

func TestDraw(t *testing.T) {
	klines := draw.Klines{Klines: []draw.KlineTOHLCV{draw.KlineTOHLCV{T: 1585564199, O: 6258.53, H: 6269.36, L: 6250.1, C: 6262.95, V: 550.501354}, draw.KlineTOHLCV{T: 1585565099, O: 6262.8, H: 6299.98, L: 6260.03, C: 6287.78, V: 1684.498345}, draw.KlineTOHLCV{T: 1585565999, O: 6288.31, H: 6313.31, L: 6261.57, C: 6277.46, V: 1787.309783}, draw.KlineTOHLCV{T: 1585566899, O: 6277.46, H: 6298, L: 6265.94, C: 6283.47, V: 925.548896}, draw.KlineTOHLCV{T: 1585567799, O: 6283.47, H: 6324.92, L: 6278.88, C: 6316.53, V: 1270.28275}, draw.KlineTOHLCV{T: 1585568699, O: 6316.55, H: 6395, L: 6316.55, C: 6350, V: 3440.725033}, draw.KlineTOHLCV{T: 1585569599, O: 6351.19, H: 6365, L: 6344.87, C: 6355.56, V: 1279.99612}, draw.KlineTOHLCV{T: 1585570499, O: 6355.55, H: 6379, L: 6316.01, C: 6322.93, V: 1531.857274}, draw.KlineTOHLCV{T: 1585571399, O: 6322.93, H: 6347.94, L: 6311, C: 6341.16, V: 1054.25279}, draw.KlineTOHLCV{T: 1585572299, O: 6341.16, H: 6341.2, L: 6288.41, C: 6325.48, V: 1724.978128}, draw.KlineTOHLCV{T: 1585573199, O: 6325.48, H: 6325.81, L: 6278.23, C: 6302.75, V: 1142.77527}, draw.KlineTOHLCV{T: 1585574099, O: 6302.76, H: 6331.99, L: 6300.44, C: 6322.69, V: 974.646619}, draw.KlineTOHLCV{T: 1585574999, O: 6322.64, H: 6344, L: 6302.58, C: 6341.32, V: 801.757113}, draw.KlineTOHLCV{T: 1585575899, O: 6342.6, H: 6346, L: 6308.89, C: 6314.55, V: 919.744509}, draw.KlineTOHLCV{T: 1585576799, O: 6314, H: 6327.09, L: 6301, C: 6302.53, V: 630.972806}, draw.KlineTOHLCV{T: 1585577699, O: 6302.64, H: 6305.95, L: 6260.62, C: 6276.81, V: 1679.465691}, draw.KlineTOHLCV{T: 1585578599, O: 6276.81, H: 6330, L: 6275.83, C: 6312.79, V: 1443.123599}, draw.KlineTOHLCV{T: 1585579499, O: 6312.92, H: 6337.42, L: 6307.56, C: 6328, V: 836.343481}, draw.KlineTOHLCV{T: 1585580399, O: 6328, H: 6332.34, L: 6316.49, C: 6326.27, V: 583.018997}, draw.KlineTOHLCV{T: 1585581299, O: 6326.27, H: 6370.74, L: 6326.27, C: 6370.01, V: 1416.269034}, draw.KlineTOHLCV{T: 1585582199, O: 6370.18, H: 6393, L: 6355.68, C: 6374.9, V: 1794.225203}, draw.KlineTOHLCV{T: 1585583099, O: 6374.91, H: 6444, L: 6346.56, C: 6365.74, V: 3495.354909}, draw.KlineTOHLCV{T: 1585583999, O: 6366.18, H: 6366.18, L: 6320, C: 6343.58, V: 2110.01719}, draw.KlineTOHLCV{T: 1585584899, O: 6343.9, H: 6352.25, L: 6303.34, C: 6338.36, V: 1547.796032}, draw.KlineTOHLCV{T: 1585585799, O: 6339.21, H: 6369.94, L: 6333.5, C: 6353.04, V: 998.005009}, draw.KlineTOHLCV{T: 1585586699, O: 6353.89, H: 6382, L: 6347.06, C: 6361.19, V: 1054.012294}, draw.KlineTOHLCV{T: 1585587599, O: 6361, H: 6389.72, L: 6353.49, C: 6377.88, V: 782.055613}, draw.KlineTOHLCV{T: 1585588499, O: 6377.86, H: 6384.98, L: 6350, C: 6359.08, V: 736.50127}, draw.KlineTOHLCV{T: 1585589399, O: 6359.07, H: 6365.17, L: 6337.82, C: 6344.85, V: 826.319426}, draw.KlineTOHLCV{T: 1585590299, O: 6344.85, H: 6354.33, L: 6330, C: 6346.99, V: 730.770487}, draw.KlineTOHLCV{T: 1585591199, O: 6347, H: 6365.11, L: 6337.45, C: 6350.46, V: 694.979102}, draw.KlineTOHLCV{T: 1585592099, O: 6350.01, H: 6354.19, L: 6325.95, C: 6335.4, V: 529.529737}, draw.KlineTOHLCV{T: 1585592999, O: 6334.62, H: 6348.61, L: 6322.49, C: 6346.37, V: 474.257231}, draw.KlineTOHLCV{T: 1585593899, O: 6346.37, H: 6374.99, L: 6345.15, C: 6365.17, V: 606.137494}, draw.KlineTOHLCV{T: 1585594799, O: 6364.23, H: 6372.03, L: 6355.87, C: 6363.01, V: 573.913118}, draw.KlineTOHLCV{T: 1585595699, O: 6362.72, H: 6372.99, L: 6350, C: 6365.97, V: 385.146807}, draw.KlineTOHLCV{T: 1585596599, O: 6365.96, H: 6402.41, L: 6348.59, C: 6357.63, V: 1053.180348}, draw.KlineTOHLCV{T: 1585597499, O: 6357.63, H: 6367.47, L: 6346.18, C: 6356.65, V: 510.66755}, draw.KlineTOHLCV{T: 1585598399, O: 6357.04, H: 6375, L: 6350, C: 6366.74, V: 502.640451}, draw.KlineTOHLCV{T: 1585599299, O: 6366.5, H: 6421.98, L: 6363.78, C: 6415.69, V: 1517.22031}, draw.KlineTOHLCV{T: 1585600199, O: 6416.79, H: 6500, L: 6416.79, C: 6438.82, V: 3744.50416}, draw.KlineTOHLCV{T: 1585601099, O: 6440.16, H: 6463.12, L: 6428.44, C: 6448.35, V: 949.826575}, draw.KlineTOHLCV{T: 1585601999, O: 6448.17, H: 6474.93, L: 6444, C: 6463.3, V: 822.017293}, draw.KlineTOHLCV{T: 1585602899, O: 6463.3, H: 6564.33, L: 6462.37, C: 6530.9, V: 3233.515709}, draw.KlineTOHLCV{T: 1585603799, O: 6532.23, H: 6534.77, L: 6477.71, C: 6506.61, V: 1299.139047}, draw.KlineTOHLCV{T: 1585604699, O: 6506.61, H: 6530.05, L: 6480.01, C: 6494.49, V: 973.602438}, draw.KlineTOHLCV{T: 1585605599, O: 6494.48, H: 6550, L: 6494.19, C: 6528.47, V: 937.839415}, draw.KlineTOHLCV{T: 1585606499, O: 6529.24, H: 6556.19, L: 6501.27, C: 6512.84, V: 1131.485494}, draw.KlineTOHLCV{T: 1585607399, O: 6512.84, H: 6590, L: 6512.83, C: 6590, V: 1024.983079}, draw.KlineTOHLCV{T: 1585608299, O: 6590, H: 6599, L: 6531.72, C: 6540, V: 2049.633016}, draw.KlineTOHLCV{T: 1585609199, O: 6540, H: 6540.55, L: 6440.29, C: 6458, V: 2389.462661}, draw.KlineTOHLCV{T: 1585610099, O: 6458.01, H: 6480, L: 6425, C: 6465.01, V: 1714.066675}, draw.KlineTOHLCV{T: 1585610999, O: 6465.01, H: 6483.54, L: 6440, C: 6459.3, V: 840.893519}, draw.KlineTOHLCV{T: 1585611899, O: 6460, H: 6462.44, L: 6411, C: 6428.31, V: 1057.20388}, draw.KlineTOHLCV{T: 1585612799, O: 6427.77, H: 6432.85, L: 6375, C: 6394.38, V: 1686.022071}, draw.KlineTOHLCV{T: 1585613699, O: 6394.45, H: 6456, L: 6380, C: 6431.41, V: 1177.121066}, draw.KlineTOHLCV{T: 1585614599, O: 6432.56, H: 6467.08, L: 6427.16, C: 6453.92, V: 746.896673}, draw.KlineTOHLCV{T: 1585615499, O: 6454.65, H: 6469.28, L: 6437, C: 6450.05, V: 673.28614}, draw.KlineTOHLCV{T: 1585616399, O: 6450.05, H: 6465.78, L: 6438, C: 6439.55, V: 859.440947}, draw.KlineTOHLCV{T: 1585617299, O: 6439.55, H: 6474.91, L: 6439.55, C: 6474.9, V: 434.331127}, draw.KlineTOHLCV{T: 1585618199, O: 6474.9, H: 6489.59, L: 6459.31, C: 6486.7, V: 636.085003}, draw.KlineTOHLCV{T: 1585619099, O: 6487.72, H: 6502.64, L: 6477.11, C: 6500.77, V: 763.21607}, draw.KlineTOHLCV{T: 1585619999, O: 6500.67, H: 6515, L: 6483.03, C: 6514.07, V: 731.181902}, draw.KlineTOHLCV{T: 1585620899, O: 6514.78, H: 6523.23, L: 6469.89, C: 6471.51, V: 909.640816}, draw.KlineTOHLCV{T: 1585621799, O: 6471.96, H: 6489.37, L: 6463.06, C: 6485.39, V: 518.467193}, draw.KlineTOHLCV{T: 1585622699, O: 6485.93, H: 6497.14, L: 6481.27, C: 6495.74, V: 513.038879}, draw.KlineTOHLCV{T: 1585623599, O: 6495.81, H: 6500, L: 6445, C: 6449.66, V: 697.098298}, draw.KlineTOHLCV{T: 1585624499, O: 6449.62, H: 6450.23, L: 6408, C: 6426.16, V: 1576.70181}, draw.KlineTOHLCV{T: 1585625399, O: 6426.51, H: 6455.45, L: 6418, C: 6452, V: 459.75737}, draw.KlineTOHLCV{T: 1585626299, O: 6452, H: 6470, L: 6447.7, C: 6465.95, V: 881.391135}, draw.KlineTOHLCV{T: 1585627199, O: 6466.76, H: 6480, L: 6436.23, C: 6439.47, V: 604.210593}, draw.KlineTOHLCV{T: 1585628099, O: 6439.47, H: 6465, L: 6415.39, C: 6432.28, V: 917.27745}, draw.KlineTOHLCV{T: 1585628999, O: 6432.2, H: 6442.01, L: 6410.63, C: 6435.16, V: 882.555957}, draw.KlineTOHLCV{T: 1585629899, O: 6435.42, H: 6436.53, L: 6384.12, C: 6404.95, V: 1146.094798}, draw.KlineTOHLCV{T: 1585630799, O: 6404.05, H: 6412.2, L: 6378, C: 6398.93, V: 909.20697}, draw.KlineTOHLCV{T: 1585631699, O: 6398.09, H: 6416.51, L: 6380.96, C: 6412.54, V: 652.475207}, draw.KlineTOHLCV{T: 1585632599, O: 6412.55, H: 6453, L: 6401, C: 6434.19, V: 661.009806}, draw.KlineTOHLCV{T: 1585633499, O: 6434.19, H: 6439.76, L: 6420.99, C: 6424.43, V: 289.407376}, draw.KlineTOHLCV{T: 1585634399, O: 6424.32, H: 6425.05, L: 6402.78, C: 6415.6, V: 604.355773}, draw.KlineTOHLCV{T: 1585635299, O: 6415.6, H: 6419.13, L: 6390, C: 6406.6, V: 574.658654}, draw.KlineTOHLCV{T: 1585636199, O: 6406.61, H: 6433.26, L: 6400.97, C: 6423.94, V: 530.247053}, draw.KlineTOHLCV{T: 1585637099, O: 6423.96, H: 6465, L: 6416.31, C: 6448.01, V: 1240.82924}, draw.KlineTOHLCV{T: 1585637999, O: 6448.01, H: 6471.56, L: 6440, C: 6471.56, V: 747.790965}, draw.KlineTOHLCV{T: 1585638899, O: 6471.56, H: 6473.68, L: 6442.73, C: 6456.97, V: 539.70587}, draw.KlineTOHLCV{T: 1585639799, O: 6457.01, H: 6459.97, L: 6446.63, C: 6456.59, V: 393.595345}, draw.KlineTOHLCV{T: 1585640699, O: 6456.59, H: 6457.42, L: 6425.48, C: 6434, V: 592.126992}, draw.KlineTOHLCV{T: 1585641599, O: 6433.1, H: 6449.08, L: 6433.09, C: 6445.09, V: 482.838777}, draw.KlineTOHLCV{T: 1585642499, O: 6445.11, H: 6465, L: 6441.71, C: 6454, V: 666.714893}, draw.KlineTOHLCV{T: 1585643399, O: 6454, H: 6488.8, L: 6451.14, C: 6469.11, V: 868.097933}, draw.KlineTOHLCV{T: 1585644299, O: 6469.68, H: 6505.5, L: 6465.5, C: 6496, V: 1148.301857}, draw.KlineTOHLCV{T: 1585645199, O: 6496.67, H: 6496.79, L: 6475, C: 6493.93, V: 729.055064}, draw.KlineTOHLCV{T: 1585646099, O: 6492.98, H: 6499, L: 6464, C: 6467.05, V: 731.717416}, draw.KlineTOHLCV{T: 1585646999, O: 6467.05, H: 6484.95, L: 6455.01, C: 6482.3, V: 592.343159}, draw.KlineTOHLCV{T: 1585647899, O: 6482.3, H: 6495, L: 6468.55, C: 6489, V: 703.872833}, draw.KlineTOHLCV{T: 1585648799, O: 6488.14, H: 6493.64, L: 6425, C: 6440.04, V: 1337.730175}, draw.KlineTOHLCV{T: 1585649699, O: 6440.04, H: 6440.04, L: 6417.97, C: 6422.74, V: 203.082677}}, Last: 6422.74, Min: 6250.1, Max: 6599, StartTime: 1.5855633e+09} //nolint
	f, err := formula.NewBasic("rate-10+0.0001*(now-start)^1.3", klines.Last, 1585649699)
	assert.Nil(t, err)
	p, err := draw.NewPlot()
	assert.Nil(t, err)
	p.DrawEnv()
	p.DrawHelpLines(klines.Last, klines.Min, klines.Max, klines.StartTime)
	err = p.DrawMainGraph(klines)
	p.DrawFunction(f, klines.Max+1.7*math.Sqrt(klines.Max))
	assert.Nil(t, err)
	err = p.Plot.Save(draw.DefaultWidth, draw.DefaultHeight, "plot.png")
	assert.Nil(t, err)
}
