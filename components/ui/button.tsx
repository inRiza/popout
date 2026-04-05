import Feather from '@expo/vector-icons/Feather';
import { TouchableOpacity, Text, StyleSheet } from 'react-native';

export default function CustomButton ({ 
    onPress, 
    title, 
    style,
    colorTitle = 'white',
    iconsName,
}: { onPress: () => void, title: string, style: any, colorTitle?: any, iconsName?: any }) {
    return (
        <TouchableOpacity onPress={onPress} style={[style, {
            borderRadius: 50,
            padding: 20,
            flexDirection: 'row',
            gap: 10,
            alignItems: 'center',
            justifyContent: 'center',
        }]}>
            <Text style={{
                color: colorTitle,
                fontSize: 16,
                fontWeight: 'bold',
            }}>{title}</Text>
            {iconsName && <Feather name={iconsName} size={18} color="white" />}
        </TouchableOpacity>
    );
}
