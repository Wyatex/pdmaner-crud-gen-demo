declare namespace Lbra {
  interface LbraCatalogAddModel {
    /** 分类名 */
    catalogName: string;
    /** 分类描述 */
    catalogDesc: string;
    /** 排序号 */
    sortNo: number;
    /** 演示时间选择 */
    demoTime: string;
  }
  interface LbraCatalogEntity {
    /** 分类ID */
    id: number;
    /** 分类名 */
    catalogName: string;
    /** 分类描述 */
    catalogDesc: string;
    /** 排序号 */
    sortNo: number;
    /** 演示时间选择 */
    demoTime: string;
    /** 创建人 */
    createdBy: string;
    /** 创建时间 */
    createdAt: string;
    /** 更新人 */
    updatedBy: string;
    /** 更新时间 */
    updatedAt: string;
    /** 删除人 */
    deletedBy: string;
    /** 删除时间 */
    deletedAt: string;
  }
  interface LbraCatalogFindPage {
    /** 分类名 */
    catalogName: string;
    /** 分类描述 */
    catalogDesc: string;
    /** 排序号 */
    sortNo: number;
    /** 演示时间选择 */
    demoTime: string;
    /** 搜索创建时间范围 */
    createdStart?: string;
    /** 搜索创建时间范围 */
    createdEnd?: string;
    /** 搜索更新时间范围 */
    updatedStart?: string;
    /** 搜索更新时间范围 */
    updatedEnd?: string;
  }
}
