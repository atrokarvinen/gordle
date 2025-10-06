import * as charts from 'echarts';

// @eslint-disable-next-line @typescript-eslint/no-explicit-any
export function echarts(node: any, option: any) {
	const chart = charts.init(node);
	chart.setOption(option);
}
