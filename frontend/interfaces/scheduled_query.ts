import PropTypes from "prop-types";
import scheduledQueryStatsInterface, {
  IScheduledQueryStats,
} from "./scheduled_query_stats";

export default PropTypes.shape({
  created_at: PropTypes.string,
  updated_at: PropTypes.string,
  id: PropTypes.number.isRequired,
  pack_id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  query_id: PropTypes.number.isRequired,
  query_name: PropTypes.string.isRequired,
  query: PropTypes.string.isRequired,
  interval: PropTypes.number.isRequired,
  snapshot: PropTypes.bool,
  removed: PropTypes.bool,
  platform: PropTypes.string,
  version: PropTypes.string,
  shard: PropTypes.number,
  denylist: PropTypes.bool,
  stats: scheduledQueryStatsInterface,
});

export interface IPackQueryFormData {
  interval: number;
  name?: string;
  shard: number;
  query?: string;
  query_id?: number;
  pack_id?: number;
  logging_type?: string;
  removed?: boolean;
  snapshot?: boolean;
  platform: string;
  version: string;
}
export interface IScheduledQuery {
  created_at: string;
  updated_at: string;
  id: number;
  pack_id: number;
  name: string;
  query_id: number;
  query_name: string;
  query: string;
  interval: number;
  snapshot: boolean;
  removed: boolean;
  platform?: string;
  version?: string;
  shard: number | null;
  denylist?: boolean;
  logging_type?: string;
  stats: IScheduledQueryStats;
}
