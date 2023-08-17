import Plan from "./plan"
import User from "./user"

type ExpandedStatistics = {
    [key: string]: any;
    user: User;
    plan: Plan;
    totalConsumptionKb: number;
    totalDailyUploadCount: number;
    totalConsumptionCount: number;
}

export default ExpandedStatistics