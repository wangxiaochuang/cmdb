package common

// GetInstNameField returns the inst name field
func GetInstNameField(objID string) string {
	switch objID {
	case BKInnerObjIDApp:
		return BKAppNameField
	case BKInnerObjIDSet:
		return BKSetNameField
	case BKInnerObjIDModule:
		return BKModuleNameField
	case BKInnerObjIDObject:
		return BKInstNameField
	case BKInnerObjIDHost:
		return BKHostNameField
	case BKInnerObjIDProc:
		return BKProcNameField
	case BKInnerObjIDPlat:
		return BKCloudNameField
	case BKTableNameInstAsst:
		return BKFieldID
	default:
		if IsObjectInstAsstShardingTable(objID) {
			return BKFieldID
		}
		return BKInstNameField
	}
}

// GetInstIDField get primary key of object's collection/table
func GetInstIDField(objType string) string {
	switch objType {
	case BKInnerObjIDApp:
		return BKAppIDField
	case BKInnerObjIDSet:
		return BKSetIDField
	case BKInnerObjIDModule:
		return BKModuleIDField
	case BKInnerObjIDObject:
		return BKInstIDField
	case BKInnerObjIDHost:
		return BKHostIDField
	case BKInnerObjIDProc:
		return BKProcIDField
	case BKInnerObjIDPlat:
		return BKCloudIDField
	case BKTableNameInstAsst:
		return BKFieldID
	case BKTableNameServiceInstance:
		return BKFieldID
	case BKTableNameServiceTemplate:
		return BKFieldID
	case BKTableNameProcessTemplate:
		return BKFieldID
	case BKTableNameProcessInstanceRelation:
		return BKProcessIDField
	default:
		if IsObjectInstAsstShardingTable(objType) {
			return BKFieldID
		}
		return BKInstIDField
	}
}

func GetObjByType(objType string) string {
	switch objType {
	case BKInnerObjIDApp, BKInnerObjIDSet,
		BKInnerObjIDModule, BKInnerObjIDProc,
		BKInnerObjIDHost, BKInnerObjIDPlat:
		return objType
	default:
		return BKInnerObjIDObject
	}
}

func IsInnerModel(objType string) bool {
	return GetObjByType(objType) != BKInnerObjIDObject
}

// IsInnerMainlineModel judge if the object type is inner mainline model
func IsInnerMainlineModel(objType string) bool {
	switch objType {
	case BKInnerObjIDApp, BKInnerObjIDSet, BKInnerObjIDModule:
		return true
	default:
		return false
	}
}
