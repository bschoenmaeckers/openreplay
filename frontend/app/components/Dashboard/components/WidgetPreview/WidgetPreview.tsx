import React from 'react';
import cn from 'classnames';
import WidgetWrapper from '../WidgetWrapper';
import { useStore } from 'App/mstore';
import { SegmentSelection } from 'UI';
import { useObserver } from 'mobx-react-lite';
import SelectDateRange from 'Shared/SelectDateRange';

interface Props {
    className?: string;
}
function WidgetPreview(props: Props) {
    const { className = '' } = props;
    const { metricStore, dashboardStore } = useStore();
    const period = useObserver(() => dashboardStore.period);
    const metric: any = useObserver(() => metricStore.instance);
    const isTimeSeries = metric.metricType === 'timeseries';
    const isTable = metric.metricType === 'table';

    const chagneViewType = (e, { name, value }: any) => {
        metric.update({ [ name ]: value });
    }

    return useObserver(() => (
        <div className={cn(className)}>
            <div className="flex items-center justify-between">
                <h2 className="text-2xl">Trend</h2>
                <div className="flex items-center">
                    {isTimeSeries && (
                        <>
                            <span className="mr-4 color-gray-medium">Visualization</span>
                            <SegmentSelection
                                name="viewType"
                                className="my-3"
                                primary
                                icons={true}
                                onSelect={ chagneViewType }
                                value={{ value: metric.viewType }}
                                list={ [
                                { value: 'lineChart', name: 'Chart', icon: 'graph-up-arrow' },
                                { value: 'progress', name: 'Progress', icon: 'hash' },
                                ]}
                            />
                        </>
                    )}

                    {isTable && (
                        <>
                            <span className="mr-4 color-gray-medium">Visualization</span>
                            <SegmentSelection
                                name="viewType"
                                className="my-3"
                                primary={true}
                                icons={true}
                                onSelect={ chagneViewType }
                                value={{ value: metric.viewType }}
                                list={[
                                { value: 'table', name: 'Table', icon: 'table' },
                                { value: 'pieChart', name: 'Chart', icon: 'pie-chart-fill' },
                                ]}
                            />
                        </>
                    )}
                    <div className="mx-4" />
                        <span className="mr-1 color-gray-medium">Time Range</span>
                        <SelectDateRange
                            period={period}
                            onChange={(period: any) => dashboardStore.setPeriod(period)}
                        />
                    </div>
            </div>
            <div className="bg-white rounded p-4">
                <WidgetWrapper widget={metric} isPreview={true} isWidget={false} />
            </div>
        </div>
    ));
}

export default WidgetPreview;